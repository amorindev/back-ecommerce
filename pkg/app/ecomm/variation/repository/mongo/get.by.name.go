package mongo

import (
	"context"
	"errors"

	errVariation "com.fernando/pkg/app/ecomm/variation/errors"
	"com.fernando/pkg/app/ecomm/variation/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetByName(ctx context.Context, name string) (*model.Variation, error) {
	var variation model.Variation

	err := r.Collection.FindOne(ctx, bson.D{{Key: "name", Value: name}}).Decode(&variation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errVariation.ErrVariationNotFound
		}
		return nil, err
	}

	objID, ok := variation.ID.(bson.ObjectID)
	if !ok {
		// pero aqui ya se creo el registro
		return nil, errors.New("variation mongo repo - GetByName, failed to parse ID to ObjectID")
	}
	variation.ID = objID.Hex()
	return &variation, nil

}
