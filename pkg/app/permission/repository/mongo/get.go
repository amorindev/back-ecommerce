package mongo

import (
	"context"
	"fmt"

	errPermission "com.fernando/pkg/app/permission/errors"
	"com.fernando/pkg/app/permission/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Get(ctx context.Context, id string) (*model.Permission, error) {
	var permission model.Permission

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("permission mongo repo - error: %v", err)
	}

	err = r.Collection.FindOne(ctx, bson.D{{Key: "_id", Value: objID}}).Decode(&permission)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errPermission.ErrPermissionNotFound
		}
		return nil, err
	}
	permission.ID = id
	return &permission, nil

}
