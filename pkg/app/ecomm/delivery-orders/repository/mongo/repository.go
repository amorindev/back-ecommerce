package mongo

import (
	"com.fernando/pkg/app/ecomm/delivery-orders/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.DeliveryOrderRepository = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewDeliveryRepo(client *mongo.Client, collection *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: collection,
	}
}
