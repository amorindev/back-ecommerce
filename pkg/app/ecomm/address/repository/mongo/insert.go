package mongo

import (
	"context"
	"errors"
	"fmt"

	"com.fernando/pkg/app/ecomm/address/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, address *model.Address) error {
	id := bson.NewObjectID()
	address.ID = id

	var userID string
	if address.UserID != nil {

		// va depender si el address pertenece a un user o un store
		userObjID, err := bson.ObjectIDFromHex(address.UserID.(string))
		if err != nil {
			return errors.New("address- failed to parse to objID")
		}

		address.UserID = userObjID
		userID = userObjID.Hex()
	}
	var storeID string
	if address.StoreID != nil {
		storeObjID, err := bson.ObjectIDFromHex(address.StoreID.(string))
		if err != nil {
			return errors.New("address- failed to parse to objID")
		}

		address.StoreID = storeObjID
		storeID = storeObjID.Hex()
	}

	_, err := r.Collection.InsertOne(context.Background(), address)
	if err != nil {
		return fmt.Errorf("address mongo repo: Insert err: %w", err)
	}
	address.ID = id.Hex()
	if address.UserID != nil {
		address.UserID = userID
	}
	if address.StoreID != nil {
		address.StoreID = storeID
	}
	//if result.InsertedID == nil ?
	return nil
}
