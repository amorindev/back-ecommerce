package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/stores/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, store *model.Store) error {
	id := bson.NewObjectID()
	store.ID = id

	_, err := r.Collection.InsertOne(context.Background(), store)
	if err != nil {
		return fmt.Errorf("store mongo repo: Insert err: %w", err)
	}
	store.ID = id.Hex()

	//if result.InsertedID == nil ?
	return nil
}
