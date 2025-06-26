package mongo

import (
	"com.fernando/pkg/app/ecomm/variation/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.VariationRepo = &Repository{}

type Repository struct{
	Client *mongo.Client
	Collection *mongo.Collection
}

func NewRepository(client *mongo.Client, variationColl *mongo.Collection) *Repository {
	return &Repository{
		Client: client,
		Collection: variationColl,
	}
}