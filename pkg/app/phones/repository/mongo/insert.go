package mongo

import (
	"context"
	"errors"
	"fmt"

	"com.fernando/pkg/app/phones/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)


func (r *Repository) Insert(ctx context.Context, phone *model.Phone) error {
	id := bson.NewObjectID()
	phone.ID = id

	userObjID, err := bson.ObjectIDFromHex(phone.UserID.(string))
	if err != nil {
		return errors.New("phone - failed to parse to objID")
	}

	phone.UserID = userObjID

	_, err = r.Collection.InsertOne(context.Background(), phone)
	if err != nil {
		return fmt.Errorf("phone mongo repo: Insert err: %w", err)
	}
	phone.ID = id.Hex()
	phone.UserID = userObjID.Hex()
	//if result.InsertedID == nil ?
	return nil
}
