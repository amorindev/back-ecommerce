package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) CreateMany(ctx context.Context, names []string) error {

	var rolesDocs []bson.D

	//  bson.D{{Key: "email", Value: email}}
	for _, r := range names {
		role := bson.D{{Key: "name", Value: r}}
		rolesDocs = append(rolesDocs, role)
	}
	_, err := r.Collection.InsertMany(context.TODO(), rolesDocs)
	if err != nil {
		return fmt.Errorf("role mongo repo - CreateMany err: %w", err)
	}

	return nil
}

/* func (r *Repository) CreateMany2(ctx context.Context, roles []string) error {

	var rolesDocs []bson.D

	//  bson.D{{Key: "email", Value: email}}
	for _, r := range roles {
		role := bson.D{{Key: "role", Value: r}}
		rolesDocs = append(rolesDocs, role)
	}
	result, err := r.Collection.InsertMany(context.TODO(), rolesDocs)
	if err != nil {
		return fmt.Errorf("role mongo repo - CreateMany err: %w", err)
	}

	objsID := result.InsertedIDs
	var IDs []string

	for _, oID := range objsID {
		idStr, ok := oID.(bson.ObjectID)
		if !ok {
			return errors.New("role mongo repo - CreateMany(number) no se puedo parser a ObjectID")
		}
		IDs = append(IDs, idStr.Hex())
	}

	return nil
} */


