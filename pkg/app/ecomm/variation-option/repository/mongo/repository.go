package mongo

import (
	"com.fernando/pkg/app/ecomm/variation-option/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.VariationOptionRepo = &Repository{}
type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewRepository(client *mongo.Client, varOptColl *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: varOptColl,
	}
}
