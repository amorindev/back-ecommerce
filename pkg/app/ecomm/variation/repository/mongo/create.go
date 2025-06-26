package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/variation/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Create(ctx context.Context, variation *model.Variation) error {
	id := bson.NewObjectID()
	variation.ID = id

	_, err := r.Collection.InsertOne(context.Background(), variation)
	if err != nil {
	  return fmt.Errorf("category mongo repo: Create err: %w", err)
	}
	variation.ID = id.Hex()
	return nil
}

