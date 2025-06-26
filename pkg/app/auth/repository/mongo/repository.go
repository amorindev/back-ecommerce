package mongo

import (
	"com.fernando/pkg/app/auth/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.AuthRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewRepository(client *mongo.Client, authColl *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: authColl,
	}
}
