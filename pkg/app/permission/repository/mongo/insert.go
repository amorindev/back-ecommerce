package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/permission/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, permission *model.Permission) error {
	id := bson.NewObjectID()
	permission.ID = id

	_, err := r.Collection.InsertOne(context.Background(), permission)
	if err != nil {
		return fmt.Errorf("permission mongo repo: Insert err: %w", err)
	}
	permission.ID = id.Hex()

	//if result.InsertedID == nil ?
	return nil
}
