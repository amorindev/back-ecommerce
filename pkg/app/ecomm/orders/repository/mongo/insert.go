package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/orders/errors"
	"com.fernando/pkg/app/ecomm/orders/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Insert(ctx context.Context, order *model.Order) error {
	id := bson.NewObjectID()
	order.ID = id

	userObjID, err := bson.ObjectIDFromHex(order.UserID.(string))
	if err != nil {
		return fmt.Errorf("order mongo repo - Insert err: %w", err)
	}
	order.UserID = userObjID
	// no quiero insertar este agregado
	orderAux := order.PaymentAgt
	order.PaymentAgt = nil

	_, err = r.Collection.InsertOne(ctx, order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.ErrOrderNotFound
		}
		return err
	}
	order.PaymentAgt = orderAux
	order.ID = id.Hex()
	order.UserID = userObjID.Hex()
	return nil
}
