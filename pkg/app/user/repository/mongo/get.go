package mongo

import (
	"context"
	"errors"

	userErr "com.fernando/pkg/app/user/errors"
	"com.fernando/pkg/app/user/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// * ver el agregate funcion de collection
func (r *Repository) Get(ctx context.Context, id string) (*model.User, error) {
	// probar si funciona solo con string
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid type ID")
	}

	var user model.User
	err = r.Collection.FindOne(ctx, bson.D{{Key: "_id", Value: objID}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, userErr.ErrUserNotFound
		}
		return nil, err
	}
	user.ID = id
	return &user, nil
}
