package mongo

import (
	"context"
	"errors"
	"fmt"

	"com.fernando/pkg/app/user/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r Repository) CreateEbd(ctx context.Context, user *model.User) error {
	id := bson.NewObjectID()
	user.ID = id

	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("auth mongo repo - Create error: %w", err)
	}

	// bueno deveria ser nil tanto para crear como para devolver
	// es por eso su existencia de esta función
	//auth.User = userAux
	user.ID = id.Hex()

	return errors.New("user mongo repo - CreateEbd review")
}
