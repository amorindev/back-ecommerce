package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/role/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, role *model.Role) error {
	role.ID = bson.NewObjectID()

	_, err := r.Collection.InsertOne(context.Background(), role)
	if err != nil {
		return fmt.Errorf("role mono repo: Create err: %w", err)
	}

	return nil
}
