package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/products/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, product *model.Product) error {
	id := bson.NewObjectID()
	product.ID = id

	ctgObjID, err := bson.ObjectIDFromHex(product.CategoryID.(string))
	if err != nil {
		return fmt.Errorf("product mongo repo - insert ObjectID from hex err %w", err)
	}

	product.CategoryID = ctgObjID
	// * pro que no te estaa saeindp
	products := product.ProductItems
	product.ProductItems = nil

	_, err = r.Collection.InsertOne(ctx, product)
	if err != nil {
		return fmt.Errorf("product mongo repo - Insert InsertOne err %w", err)
	}
	product.ID = id.Hex()
	product.CategoryID = ctgObjID.Hex()

	product.ProductItems = products

	// * pasar el id a los productos o desde donde hacerlo?
	for _, product := range product.ProductItems {
		product.ProductID = id.Hex()
	}

	return nil
}
