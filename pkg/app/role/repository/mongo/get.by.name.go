package mongo

import (
	"context"
	"fmt"

	roleErr "com.fernando/pkg/app/role/errors"
	"com.fernando/pkg/app/role/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// !hacer elfor desde el repo como insert []roles debe tener transaccion?
// o hacer el for desde fuera
func (r *Repository) GetByName(ctx context.Context, name string) (*model.Role, error) {
	var role model.Role

	// ? se puede retornar solo el campo id y no todo el objeto ?
	err := r.Collection.FindOne(ctx, bson.D{{Key: "name", Value: name}}).Decode(&role)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, roleErr.ErrRoleNotFound
		}
		return nil, fmt.Errorf("role mongo repo - GetByName err: %w", err)
	}

	objID, ok := role.ID.(bson.ObjectID)
	if !ok {
		return nil, fmt.Errorf("role mongo repo -GetByName, failed to convert ID to ObjectID")
	}
	role.ID = objID.Hex()

	return &role, nil
}
