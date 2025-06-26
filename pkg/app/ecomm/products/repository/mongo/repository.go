package mongo

import (
	"com.fernando/pkg/app/ecomm/products/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.ProductRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewRepository(client *mongo.Client, productColl *mongo.Collection) *Repository {
	return &Repository{
		Client: client, Collection: productColl,
	}
}
