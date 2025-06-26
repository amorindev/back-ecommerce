package mongo

import (
	"context"
	"fmt"

	authErr "com.fernando/pkg/app/auth/errors"
	"com.fernando/pkg/app/auth/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetByIDProvider(ctx context.Context, userID string, provider string) (*model.Auth, error) {
	var auth model.Auth

	objID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("auth mongo repo - GetByIDProvider err: %w", err)
	}

	filter := bson.M{"user_id": objID, "provider": provider}

	err = r.Collection.FindOne(ctx, filter).Decode(&auth)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, authErr.ErrAuthNotFound
		}
		return nil, fmt.Errorf("auth mongo repo - GetByIDProvider err: %w", err)
	}

	auth.ID = userID

	return &auth, nil
}
