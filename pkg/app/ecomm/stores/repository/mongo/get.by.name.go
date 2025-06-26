package mongo

import (
	"context"
	"errors"

	storeErr "com.fernando/pkg/app/ecomm/stores/errors"
	"com.fernando/pkg/app/ecomm/stores/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetByName(ctx context.Context, name string) (*model.Store, error) {
	var store model.Store

	err := r.Collection.FindOne(ctx, bson.D{{Key: "name", Value: name}}).Decode(&store)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, storeErr.ErrStoreNotFound
		}
		return nil, err
	}
	objID, ok := store.ID.(bson.ObjectID)
	if !ok {
		// pero aqui ya se creo el registro
		return nil, errors.New("store mongo repo - GetByName, failed to parse ID to ObjectID")
	}
	store.ID = objID.Hex()
	return &store, nil
}
