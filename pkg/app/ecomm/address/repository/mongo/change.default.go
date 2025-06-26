package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) ChangeDefault(ctx context.Context, id string, isDefault bool) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"is_default": isDefault,
		},
	}

	_, err = r.Collection.UpdateOne(ctx, filter, update)
	return err
}
