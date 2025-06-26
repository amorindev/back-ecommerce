package mongo

import (
	"com.fernando/pkg/app/ecomm/pickup-orders/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.PickupOrderRepository = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewPickupRepo(client *mongo.Client, collection *mongo.Collection) *Repository{
	return &Repository{
		Client: client,
		Collection: collection,
	}
}
