package mongo

import (
	"context"
	"errors"

	errCategory "com.fernando/pkg/app/ecomm/category/errors"
	"com.fernando/pkg/app/ecomm/category/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetByName(ctx context.Context, name string) (*model.Category, error) {
	var category model.Category

	err := r.Collection.FindOne(ctx, bson.D{{Key: "name", Value: name}}).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errCategory.ErrCategoryNotFound
		}
		return nil, err
	}

	objID, ok := category.ID.(bson.ObjectID)
	if !ok {
		// pero aqui ya se creo el registro
		return nil, errors.New("category mongo repo - GetByName, failed to parse ID to ObjectID")
	}
	category.ID = objID.Hex()
	return &category, nil
}
