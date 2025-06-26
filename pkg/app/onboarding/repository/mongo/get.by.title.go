package mongo

import (
	"context"
	"errors"

	errOnboarding "com.fernando/pkg/app/onboarding/errors"
	"com.fernando/pkg/app/onboarding/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetByTitle(ctx context.Context, title string) (*model.Onboarding, error) {
	var onboarding model.Onboarding

	err := r.Collection.FindOne(ctx, bson.D{{Key: "title", Value: title}}).Decode(&onboarding)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errOnboarding.ErrOnboardingNotFound
		}
		return nil, err
	}

	objID, ok := onboarding.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("onboarding mongo repo - GetByTitle, failed to parse ID to ObjectID")
	}
	onboarding.ID = objID.Hex()
	return &onboarding, nil
}
