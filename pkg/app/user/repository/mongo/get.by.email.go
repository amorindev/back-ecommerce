package mongo

import (
	"context"
	"errors"

	userErr "com.fernando/pkg/app/user/errors"
	"com.fernando/pkg/app/user/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := r.Collection.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, userErr.ErrUserNotFound
		}
		// verificar que todos retornen esto
		return nil, err
	}

	objID, ok := user.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("user mongo repo - GetByEmail, failed to parse ID to ObjectID")
	}
	user.ID = objID.Hex()

	return &user, nil
}
