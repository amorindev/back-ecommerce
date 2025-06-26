package mongo

import (
	"context"
	"errors"

	addressErr "com.fernando/pkg/app/ecomm/address/errors"
	"com.fernando/pkg/app/ecomm/address/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetDefault(ctx context.Context) (*model.Address, error) {
	var address model.Address

	err := r.Collection.FindOne(ctx, bson.D{{Key: "is_default", Value: true}}).Decode(&address)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, addressErr.ErrAddressNotFound
		}
		// verificar que todos retornen esto
		return nil, err
	}

	objID, ok := address.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("address mongo repo - GetDefault, failed to parse ID to ObjectID")
	}
	address.ID = objID.Hex()

	userObjID, ok := address.UserID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("address mongo repo - GetDefault, failed to parse ID to ObjectID")
	}
	address.UserID = userObjID.Hex()
	return &address, nil
}
