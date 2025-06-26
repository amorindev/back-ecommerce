package mongo

import (
	"com.fernando/pkg/app/ecomm/product-config/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.ProductConfigRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewRepository(client *mongo.Client, productConfigColl *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: productConfigColl,
	}
}
