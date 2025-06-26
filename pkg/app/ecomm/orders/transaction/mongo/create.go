package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/orders/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (t *Transaction) Create(ctx context.Context, order *model.Order) error {
	session, err := t.Client.StartSession()
	if err != nil {
		return fmt.Errorf("orders mongo tx - Create err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("orders mongo tx - Create err: %w", err)
	}

	defer session.EndSession(ctx)

	// ! no se si llamarlo context 2 para diferenciarlo o normal no colisiona
	err = mongo.WithSession(ctx, session, func(ctx context.Context) error {
		// ver el agregado
		err = t.OrderRepo.Insert(ctx, order)
		if err != nil {
			return err
		}

		// * para sapar el id puede ser
		// * 1. pasarlo desde el Insert (orderRepo) lo cual rompe single responsability
		// * 2. pasarlo desde la tx  order.ItemID=order.ID algo asi pero que pasa si es
		// * 	una lista (nada mal cuando no es lista)
		// * 3. (mas conveniente) InserOrderProduct agregar parámetro orderID y dentro de
		// * 	la funcion asignamos e insertamos,
		// * se podira hacer un estandar si es uno la segunda si es lista la 3ra
		// * o sino para cualquier caso la 3 unicamente
		// * otra cosa es que si viene de la re quest el id o si en el handler lo asigno
		// *

		// se necesitaria el userID
		// por que el product id ya viene desde el frontend
		// ver lo que ya viene, y ver si desde el servicio se hae eso
		// ver si se asigna desde el servicio o desde el tx
		// tambien ver que se va agregar a la funcion userID alINserOrder....
		err = t.OrderRepo.InsertOrderProduct(ctx, order.ID.(string), order.Items)
		if err != nil {
			return err
		}

		// ! primera forma abajo la segunda
		err = t.PaymentRepo.Insert(ctx, order.ID.(string), order.PaymentAgt)
		if err != nil {
			return err
		}

		if order.DeliveryType == "delivery" {
			order.DeliveryAgt.OrderID = order.ID
			err = t.DeliveryRepo.Insert(ctx, order.DeliveryAgt)
			if err != nil {
				return err
			}
		} else if order.DeliveryType == "pickup" {
			order.PickupAgt.OrderID = order.ID
			err = t.PickupRepo.Insert(ctx, order.PickupAgt)
			if err != nil {
				return err
			}
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		if err := session.AbortTransaction(context.Background()); err != nil {
			return fmt.Errorf("orders mongo transaction - Create AbortTransaction err: %w", err)
		}
		return fmt.Errorf("orders mongo transaction - Create err: %w", err)
	}
	return nil
}
