package mongo

import (
	"context"
	"errors"
	"fmt"

	"com.fernando/pkg/app/auth/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// create embedding
// ! todos los Create deben retornar el user ?
func (r *Repository) CreateEbd(ctx context.Context, auth *model.Auth) error {
	auth.ID = bson.NewObjectID()

	// ? hacerlodesde aqui o desde user repo
	// auth.ID = auth.UserAgregate.ID = id desed qui o el rtransaction?

	// por que lo pase a string, por qu se que viene string´
	// por que en get lo pase a string
	userObjID, err := bson.ObjectIDFromHex(auth.UserID.(string))
	if err != nil {
		return fmt.Errorf("auth mongo repo - CreateEbd err: %w", err)
	}
	auth.UserID = userObjID

	_, err = r.Collection.InsertOne(context.Background(), auth)
	if err != nil {
		return fmt.Errorf("auth mongo repo - CreateEbd error: %w", err)
	}

	/* idStr, ok := auth.ID.(string)
	if !ok {
		return errors.New("auth mongo repo - CreateEbd, failed to parse .(string)")
	}
	auth.ID = idStr */

	/* auth.ID = auth.ID.(string) // error */

	authObjID, ok := auth.ID.(bson.ObjectID)
	if !ok {
		return errors.New("auth mongo repo - CreateEbd, failed to convert authID to ObjectID")
	}
	auth.ID = authObjID.Hex()

	return nil
}
