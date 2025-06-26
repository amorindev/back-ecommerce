package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/onboarding/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, onboarding *model.Onboarding) error {
	id := bson.NewObjectID()
	onboarding.ID = id

	_, err := r.Collection.InsertOne(context.Background(), onboarding)
	if err != nil {
		return fmt.Errorf("onboarding mongo repo: Insert err: %w", err)
	}
	onboarding.ID = id.Hex()

	//if result.InsertedID == nil ?
	return nil
}
