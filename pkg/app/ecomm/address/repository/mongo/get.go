package mongo

import (
	"context"
	"errors"
	"fmt"

	addressErr "com.fernando/pkg/app/ecomm/address/errors"
	"com.fernando/pkg/app/ecomm/address/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Get(ctx context.Context, id string) (*model.Address, error) {
	var address model.Address

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("address mongo repo - Get err d: %w", err)
	}

	// * si le paso el user id del claim no sería mas rápido
	filter := bson.M{"_id": objID}

	err = r.Collection.FindOne(ctx, filter).Decode(&address)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, addressErr.ErrAddressNotFound
		}
		return nil, fmt.Errorf("address mongo repo - Get err: %w", err)
	}

	address.ID = objID.Hex()

	addressObjID, ok := address.UserID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("adddress mongo repo - Get, failed to convert authID to ObjectID")
	}
	address.UserID = addressObjID.Hex()

	return &address, nil
}
