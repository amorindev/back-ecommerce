package mongo

import (
	"context"

	"com.fernando/pkg/app/auth/model"
)

// TODO: No esta implementado en los puertos
// * no deberia pero para fines prácticos

// !crear la transaccion?
// como mantener mongo con embedding
func (r Repository) CreateWithUserEbd(ctx context.Context, auth *model.Auth) error {
	/* id := bson.NewObjectID()
	auth.ID = id
	auth.UserAgregate.ID = id

	session, err := r.Client.StartSession()
	if err != nil {
		return fmt.Errorf("auth mongo repo - Create: %w", err)
	}

	if err = session.StartTransaction(); err != nil {
		return fmt.Errorf("auth mongo repo - Create: %w", err)
	}


	if err = mongo.WithSession(ctx, session, func(ctx context.Context) error {

		return nil
	}); err != nil {
		return fmt.Errorf("auth mongo repo - Create: %w", err)
	}

	_, err = r.UserCollection.InsertOne(context.Background(), auth.UserAgregate)
	if err != nil {
		return fmt.Errorf("auth mongo repo - Create error: %w", err)
	}

	// ! revisar
	// si auth.User = nil
	// afectará a la capa de servicio por que necesitamos el ID del usuario
	// ademas estoy usando embedding asi que hare nosql para mantener los datos
	_, err = r.AuthCollection.InsertOne(context.Background(), auth)
	if err != nil {
		return fmt.Errorf("auth mongo repo - Create error: %w", err)
	}
	*/
	return nil
}
