package mongo

import (
	"context"
	"errors"
	"fmt"

	"com.fernando/pkg/app/role/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// ? se deberia cparsear  a string? el ID
func (r *Repository) FindByNames(ctx context.Context, names []string) (roleIDs []string, err error) {

	var roles []string

	for _, roleName := range names {
		var role model.Role
		// ? se puede retornar solo el campo id y no todo el objeto ?
		err := r.Collection.FindOne(context.TODO(), bson.D{{Key: "name", Value: roleName}}).Decode(&role)
		if err != nil {
			return nil, fmt.Errorf("role mongo repo - FindByNames error: %w", err)
		}
		id, ok := role.ID.(bson.ObjectID)
		if !ok {
			return nil, errors.New("role mongo repo - FindByNames failed to convert auth ID to string")
		}
		role.ID = id.Hex()
		roles = append(roles, id.Hex())
	}

	return roles, nil
}
