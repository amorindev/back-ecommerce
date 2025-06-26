package mongo

import (
	"context"
	"errors"

	ErrProduct "com.fernando/pkg/app/ecomm/products/errors"
	"com.fernando/pkg/app/ecomm/products/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetByName(ctx context.Context, name string) (*model.Product, error) {
	var product model.Product

	err := r.Collection.FindOne(ctx, bson.D{{Key: "name", Value: name}}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrProduct.ErrProductNotFound
		}
		println(err)

		return nil, err
	}

	objID, ok := product.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("product mongo repo - GetByName, failed to parse ID to ObjectID")
	}
	product.ID = objID.Hex()
	// TODO falta parsear los demas ids ver asi para los demas son como 4 calecciones ver todos los get
	ctgObjID, ok := product.CategoryID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("product mongo repo - GetByName, failed to parse categoryID to ObjectID")
	}
	// TODO: como hacer con los demas campos
	product.CategoryID = ctgObjID.Hex()
	return &product, nil
}
