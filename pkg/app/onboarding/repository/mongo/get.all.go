package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/onboarding/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Get(ctx context.Context) ([]*model.Onboarding, error) {
	var onboardings []*model.Onboarding

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("onboarding mongo repo- Get: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		// ? pasarlo a stirng antes de retornarlo el id o las lista mantenorlo con objeID
		// ? de momento no tube que asignar como string en el token era el inconveniente ver
		var onboarding model.Onboarding
		if err := cursor.Decode(&onboarding); err != nil {
			return nil, fmt.Errorf("onboarding mongo repo- Get: %v", err)
		}
		onboardings = append(onboardings, &onboarding)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("onboarding mongo repo- Get: %v", err)
	}
	return onboardings, nil
}
