package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/permission/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// * ver que se va hacer en cada caso ( si ya existe por  lo menos uno)solo pasarlo o cuando es critico que
// * retirne el error como crear kardex
func (r *Repository) ExistOne(ctx context.Context, permissions []*model.Permission) ([]*model.Permission, error) {
	values := make([]string, 0, len(permissions))
	for _, v := range permissions {
		// ? deepseek esta revisando nulos cuando acerlo al usar?
		if v.Name != nil {
			values = append(values, *v.Name)
		}
	}

	filter := bson.M{"name": bson.M{"$in": values}}
	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("permision error finding  options: %w", err)
	}
	defer cursor.Close(ctx)

	var existingPermissions []*model.Permission
	if err := cursor.All(ctx, &existingPermissions); err != nil {
		return nil, fmt.Errorf("permission error decoding  options: %v", err)
	}
	// * otra forma seria si existingOptions > 0 retornar false

	return existingPermissions, nil
}
