package mongo

import (
	"com.fernando/pkg/app/onboarding/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.OnboardingRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewRepository(client *mongo.Client, collection *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: collection,
	}
}
