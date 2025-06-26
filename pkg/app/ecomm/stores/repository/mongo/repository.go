package mongo

import (
	"com.fernando/pkg/app/ecomm/stores/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.StoreRepo = &Repository{}

type Repository struct{
	Client *mongo.Client
	Collection *mongo.Collection
}

func NewRepository(client *mongo.Client, collection *mongo.Collection) *Repository{
	return &Repository{
		Client: client,
		Collection: collection,
	}
}