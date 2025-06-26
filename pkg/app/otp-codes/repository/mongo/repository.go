package mongo

import (
	"com.fernando/pkg/app/otp-codes/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.OtpRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewRepository(client *mongo.Client, otpColl *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: otpColl,
	}
}
