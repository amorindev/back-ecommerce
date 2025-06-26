package mongo

import (
	"com.fernando/pkg/app/role/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.RoleRepo = &Repository{}

type Repository struct {
	Client       *mongo.Client
	Collection   *mongo.Collection
	UserRoleColl *mongo.Collection
}

func NewRepository(client *mongo.Client, collection *mongo.Collection, userRoleColl *mongo.Collection) *Repository {
	return &Repository{
		Client:       client,
		Collection:   collection,
		UserRoleColl: userRoleColl,
	}
}
