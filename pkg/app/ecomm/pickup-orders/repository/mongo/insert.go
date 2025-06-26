package mongo

import (
	"context"
	"errors"
	"fmt"

	"com.fernando/pkg/app/ecomm/pickup-orders/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, pickupO *model.PickupOrder) error {
	id := bson.NewObjectID()
	pickupO.ID = id

	orderObjID, err := bson.ObjectIDFromHex(pickupO.OrderID.(string))
	if err != nil {
		return errors.New("delivery- failed to parse to objID")
	}

	phoneObjID, err := bson.ObjectIDFromHex(pickupO.PhoneID.(string))
	if err != nil {
		return errors.New("delivery- failed to parse to objID")
	}
	addressObjID, err := bson.ObjectIDFromHex(pickupO.AddressID.(string))
	if err != nil {
		return errors.New("delivery- failed to parse to objID")
	}
	storeObjID, err := bson.ObjectIDFromHex(pickupO.StoreID.(string))
	if err != nil {
		return errors.New("delivery- failed to parse to objID")
	}
	pickupO.OrderID = orderObjID
	pickupO.AddressID = addressObjID
	pickupO.PhoneID = phoneObjID
	pickupO.StoreID = storeObjID

	_, err = r.Collection.InsertOne(context.Background(), pickupO)
	if err != nil {
		return fmt.Errorf("delivery mongo repo: Insert err: %w", err)
	}

	pickupO.ID = id.Hex()
	pickupO.OrderID = orderObjID.Hex()
	pickupO.AddressID = addressObjID.Hex()
	pickupO.PhoneID = phoneObjID.Hex()
	pickupO.StoreID = storeObjID.Hex()
	//if result.InsertedID == nil ?
	return nil
}
