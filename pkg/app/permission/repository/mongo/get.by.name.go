package mongo

import (
	"context"
	"errors"

	errPermission "com.fernando/pkg/app/permission/errors"
	"com.fernando/pkg/app/permission/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetByName(ctx context.Context, name string) (*model.Permission, error) {
	var permission model.Permission

	err := r.Collection.FindOne(ctx, bson.D{{Key: "name", Value: name}}).Decode(&permission)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errPermission.ErrPermissionNotFound
		}
		return nil, err
	}
	objID, ok := permission.ID.(bson.ObjectID)
	if !ok {
		// pero aqui ya se creo el registro
		return nil, errors.New("permission mongo repo - GetByName, failed to parse ID to ObjectID")
	}
	permission.ID = objID.Hex()
	return &permission, nil

}
