package mongo

import (
	"context"
	"errors"
	"fmt"

	errPhone "com.fernando/pkg/app/phones/errors"
	"com.fernando/pkg/app/phones/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Get(ctx context.Context, id string) (*model.Phone, error) {
	//fmt.Printf("Phone ID: %s\n", id)
	var phone model.Phone

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("phone mongo repo - Get err d: %w", err)
	}

	// * si le paso el user id del claim no sería mas rápido
	filter := bson.M{"_id": objID}

	err = r.Collection.FindOne(ctx, filter).Decode(&phone)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errPhone.ErrPhoneNotFound
		}
		return nil, fmt.Errorf("phone mongo repo - Get err: %w", err)
	}

	phone.ID = objID.Hex()

	userObjID, ok := phone.UserID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("phone mongo repo - Get, failed to convert authID to ObjectID")
	}
	phone.UserID = userObjID.Hex()

	return &phone, nil
}
