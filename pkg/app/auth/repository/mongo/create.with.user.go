package mongo

import (
	"context"

	"com.fernando/pkg/app/auth/model"
)

// TODO: No esta implementado en los puertos
// Dos docs sin embed
func (r Repository) CreateWithUser(ctx context.Context, auth *model.Auth) error {
	/* id := bson.NewObjectID()
	auth.ID = id
	auth.UserAgregate.ID = id
	// se podria usar omitempty pero no bson:"-", por que afectaria a as funciones con embedding
	userAux := auth.UserAgregate

	_, err := r.UserCollection.InsertOne(context.Background(), auth.UserAgregate)
	if err != nil {
		return fmt.Errorf("auth mongo repo - Create error: %w", err)
	}

	auth.UserAgregate= nil
	_, err = r.AuthCollection.InsertOne(context.Background(), auth)
	if err != nil {
		return fmt.Errorf("auth mongo repo - Create error: %w", err)
	}

	auth.UserAgregate = userAux */
	return nil
}
