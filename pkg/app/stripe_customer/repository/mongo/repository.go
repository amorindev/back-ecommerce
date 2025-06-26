package mongo

import (
	"com.fernando/pkg/app/stripe_customer/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.StripeCustomerRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

// ! deberia ser  solo de stripe?
func NewPaymentProviderRepo(client *mongo.Client, collection *mongo.Collection) *Repository{
	return &Repository{
		Client: client,
		Collection: collection,
	}
}
