package mongo

import (
	"com.fernando/pkg/app/phones/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.PhoneRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewRepository(client *mongo.Client, coll *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: coll,
	}
}
