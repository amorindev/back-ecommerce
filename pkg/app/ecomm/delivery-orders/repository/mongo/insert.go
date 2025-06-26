package mongo

import (
	"context"
	"errors"
	"fmt"

	"com.fernando/pkg/app/ecomm/delivery-orders/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, deliveryO *model.DeliveryOrder) error {
	id := bson.NewObjectID()
	deliveryO.ID = id

	orderObjID, err := bson.ObjectIDFromHex(deliveryO.OrderID.(string))
	if err != nil {
		return errors.New("delivery- failed to parse to objID")
	}

	phoneObjID, err := bson.ObjectIDFromHex(deliveryO.PhoneID.(string))
	if err != nil {
		return errors.New("delivery- failed to parse to objID")
	}

	addressObjID, err := bson.ObjectIDFromHex(deliveryO.AddressID.(string))
	if err != nil {
		return errors.New("delivery- failed to parse to objID")
	}

	deliveryO.OrderID = orderObjID
	deliveryO.AddressID = addressObjID
	deliveryO.PhoneID = phoneObjID

	_, err = r.Collection.InsertOne(context.Background(), deliveryO)
	if err != nil {
		return fmt.Errorf("delivery mongo repo: Insert err: %w", err)
	}

	deliveryO.ID = id.Hex()
	deliveryO.OrderID = orderObjID.Hex()
	deliveryO.AddressID = addressObjID.Hex()
	deliveryO.PhoneID = phoneObjID.Hex()
	//if result.InsertedID == nil ?

	return nil
}
