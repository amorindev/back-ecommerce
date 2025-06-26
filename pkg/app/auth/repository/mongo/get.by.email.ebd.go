package mongo

import (
	"context"

	"com.fernando/pkg/app/auth/model"
)

// Si es ebd tiene que comvertir de ambos tanto de uth como user, todo el arbol de objetos id a string
// y verificar la relación hora es unp amuchos auth
/*
definir errores
var (
	ErrUserNotFound = errors.New("user-not-found") //igual a firebase
)

*/
func (r *Repository) GetByEmailEbd(ctx context.Context, email string) (*model.Auth, error) {
	/* var auth model.Auth

	err := r.AuthCollection.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&auth)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, authErr.ErrAuthNotFound
		}
		return nil, err
	}

	objID, ok := auth.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("auth mongo repo - GetByEmail failed to parse ID to ObjectID")
	}
	auth.ID = objID.Hex()

	userID, ok := auth.UserAgregate.ID.(bson.ObjectID)
	if !ok {
		return nil, fmt.Errorf("auth mongo repo - failed to convert user ID to string: %v", auth.UserAgregate.ID)
	}
	auth.UserAgregate.ID = userID.Hex()

	return &auth, nil */
	return nil, nil
}

func (r *Repository) GetByEmailEbd2(ctx context.Context, email string) (*model.Auth, error) {
	// * Dos formas: si no esta embeding neceseitara auth el email, si no buscar por embeding
	// ! separar en dos tablas user y auth, para no hacer join e email estará en ambos
	// * si el usario ingresa a la pantalla de profile ya no sería necesario ir a auth handlers
	// * ano ser que se use shared preferences pero data, o el de auth no lo se
	/* var auth model.Auth
	// ? var user model.User ; auth.User = user es necesario o cuando no es embedding
	//var result bson.M
	err := r.AuthCollection.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&auth)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, authErr.ErrAuthNotFound
		}
		return nil, fmt.Errorf("auth mogo repo - GetByEmail error: %w", err)
	}

	// converitr a string hex los IDs auth y user
	authID, ok := auth.ID.(bson.ObjectID)
	if !ok {
		return nil, fmt.Errorf("auth mongo repo - failed to convert auth ID to string: %v", auth.ID)
	}
	auth.ID = authID.Hex()

	// * si uso sol create en AuthRepo me dara error en signin
	userID, ok := auth.UserAgregate.ID.(bson.ObjectID)
	if !ok {
		return nil, fmt.Errorf("auth mongo repo - failed to convert user ID to string: %v", auth.UserAgregate.ID)
	}
	auth.UserAgregate.ID = userID.Hex() */

	/* fmt.Printf("Auth %v\n", auth)
	fmt.Printf("User %v\n", auth.User)
	if auth.User == nil {
		return nil, errors.New("auth mongo repo - user is nil")
	} */

	//return &auth, nil
	return nil, nil
}
