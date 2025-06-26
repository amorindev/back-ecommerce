package mongo

import (
	"com.fernando/pkg/app/ecomm/product-item/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.ProductItemRepo = &Repository{}

type Repository struct{
	Client *mongo.Client
	Collection *mongo.Collection
}

func NewRepository(client *mongo.Client, productItemColl *mongo.Collection) *Repository{
	return &Repository{
		Client: client,
		Collection: productItemColl,
	}
}