package mongo

import (
	"com.fernando/pkg/app/permission/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.PermissinRepo = &Repository{}

type Repository struct {
	Client             *mongo.Client
	Collection         *mongo.Collection
	RolePermissionColl *mongo.Collection
}

func NewRepository(client *mongo.Client, collection *mongo.Collection, rolePermissionColl *mongo.Collection) *Repository {
	return &Repository{
		Client:             client,
		Collection:         collection,
		RolePermissionColl: rolePermissionColl,
	}
}
