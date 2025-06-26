package mongo

import (
	"com.fernando/pkg/app/ecomm/payment/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.PaymentRepository = &Repository{}

type Repository struct{
	Client *mongo.Client
	Collection *mongo.Collection
}

func NewRepository(client *mongo.Client, coll *mongo.Collection) *Repository{
	return &Repository{
		Client: client,
		Collection: coll,
	}
}