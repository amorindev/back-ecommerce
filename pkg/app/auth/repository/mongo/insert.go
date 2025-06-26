package mongo

import (
	"context"
	"errors"
	"fmt"

	authModel "com.fernando/pkg/app/auth/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// ! Eliminar los create.with.user.*
// * solo debe estar auth

// * verificar que todos los create verifique si se a creado error no rows
func (r *Repository) Insert(ctx context.Context, auth *authModel.Auth) error {
	id := bson.NewObjectID()
	auth.ID = id

	// * obtener el ID auth y pasarle al UserID
	userObjID, err := bson.ObjectIDFromHex(auth.UserID.(string))
	if err != nil {
		return fmt.Errorf("auth mongo repo - Create err: %w", err)
	}

	auth.UserID = userObjID

	_, err = r.Collection.InsertOne(ctx, auth)
	if err != nil {
		return fmt.Errorf("auth mongo repo - Create error: %w", err)
	}

	userID, ok := auth.UserID.(bson.ObjectID)
	if !ok {
		return errors.New("auth mongo repo - Create, failed to parse ID to ObjectId")
	}

	auth.UserID = userID.Hex()
	auth.ID = id.Hex()

	return nil
}

/*
func (r *Repository) CreateForma1(ctx context.Context, auth *model.Auth) error{
	auth.ID = bson.NewObjectID()

	_, err := r.Collection.InsertOne(ctx, auth)
	if err != nil {
		return err
	}

	id, ok := auth.ID.(bson.ObjectID)
	if !ok {
		return errors.New("error al convertir a primitive object Create auth")
	}
	auth.ID = id.Hex()

	return nil
}

func (r *Repository) CreateForma2(ctx context.Context, auth *model.Auth) error {
	id := bson.NewObjectID()
	auth.ID = id

	_, err := r.Collection.InsertOne(ctx, auth)
	if err != nil {
		return err
	}

	auth.ID = id.Hex()

	return nil
} */
