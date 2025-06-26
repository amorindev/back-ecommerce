package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/orders/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// * tenemos dos comparacioenes uas bson o una estructura y verificar si existe la relacion
// * entre estas dos funciones AssignRolesToUser
// * o creamos el modelo OrderItem o usamos un bson como en AssignRolesToUser ver, comparar
func (r *Repository) InsertOrderProduct(ctx context.Context, orderID string, items []*model.OrderItem) error {
	orderObjID, err := bson.ObjectIDFromHex(orderID)
	if err != nil {
		return fmt.Errorf("order mongo repo - InserOrderProduct err: %w", err)
	}

	for _, item := range items {
		// ? varificar si existe la relación, comparar con AssignRolesToUser ahi hay un ejemplo
		id := bson.NewObjectID()
		item.ID = id
		item.OrderID = orderObjID

		//fmt.Printf("ID: %v\n", item.ProductItemID)

		productObjID, err := bson.ObjectIDFromHex(item.ProductItemID.(string))
		if err != nil {
			return fmt.Errorf("order mongo repo - InserOrderProduct err: %w", err)
		}

		item.ProductItemID = productObjID
	}

	session, err := r.Client.StartSession()
	if err != nil {
		return fmt.Errorf("order mongo repo - InserOrderProduct err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("order mongo repo - InserOrderProduct err: %w", err)
	}

	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(ctx2 context.Context) error {
		_, err = r.OrderProductColl.InsertMany(ctx2, items)
		if err != nil {
			return err
		}
		// * ver como mescalar los cotnextos
		return session.CommitTransaction(ctx2)
	})
	if err != nil {
		if err := session.AbortTransaction(context.Background()); err != nil {
			return fmt.Errorf("order mongo repo - InserOrderProduct abort transaction err: %w", err)
		}

		return fmt.Errorf("order mongo repo - InserOrderProduct err: %w", err)
	}
	return nil
}
