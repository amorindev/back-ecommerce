package mongo

import (
	"context"
	"errors"

	ErrVarOpt "com.fernando/pkg/app/ecomm/variation-option/errors"
	"com.fernando/pkg/app/ecomm/variation-option/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetByName(ctx context.Context, value string) (*model.VariationOption, error) {
	var varOption model.VariationOption

	err := r.Collection.FindOne(ctx, bson.D{{Key: "value", Value: value}}).Decode(&varOption)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrVarOpt.ErrVarOptNotFound
		}
		return nil, err
	}

	objID, ok := varOption.ID.(bson.ObjectID)
	if !ok {
		// pero aqui ya se creo el registro
		return nil, errors.New("variation option mongo repo - GetByName, failed to parse ID to ObjectID")
	}
	variationObjID, ok := varOption.VariationID.(bson.ObjectID)
	if !ok {
		// pero aqui ya se creo el registro
		return nil, errors.New("variation option mongo repo - GetByName, failed to parse ID to ObjectID")
	}

	varOption.ID = objID.Hex()
	varOption.VariationID = variationObjID.Hex()

	return &varOption, nil
}
