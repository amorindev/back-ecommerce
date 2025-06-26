package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/category/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, category *model.Category) error{
	id := bson.NewObjectID()
	category.ID = id

	_, err := r.Collection.InsertOne(context.Background(), category)
	if err != nil {
	  return fmt.Errorf("category mongo repo: Create err: %w", err)
	}
	category.ID = id.Hex()

	//if result.InsertedID == nil ?
	return nil
}

