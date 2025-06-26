package mongo

import (
	"context"
	"errors"
	"fmt"

	"com.fernando/pkg/app/user/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Create de auth service guía
func (r *Repository) Insert(ctx context.Context, user *model.User) error {
	// * nuevamente ver las validaciones que se hacen en cada capa y los punteros
	if user == nil {
		return errors.New("user is nil")
	}
	// ! Asi para todos los que usan transacción

	user.ID = bson.NewObjectID()

	aux := user.Roles
	user.Roles = nil

	_, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("user mongo repo - Create error: %w", err)
	}

	// falta convertir a string

	objID, ok := user.ID.(bson.ObjectID)
	if !ok {
		return errors.New("user  mongo repo - Create, failed to parse ID to ObjectID")
	}
	user.ID = objID.Hex()
	// * version con [0]user.AuthProviders[0].UserID = objID.Hex()
	user.AuthProviderCreate.UserID = objID.Hex()

	// lo hago por que primero se crea el user despues los roles si no quedará nil
	user.Roles = aux

	return nil
}
