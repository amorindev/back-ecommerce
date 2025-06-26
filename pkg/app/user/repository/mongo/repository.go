package mongo

import (
	"com.fernando/pkg/app/user/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.UserRepo = &Repository{}

type Repository struct {
	Client             *mongo.Client
	Collection         *mongo.Collection
	TwoFaSmsColl       *mongo.Collection
	UserRoleCollection *mongo.Collection
}

func NewRepository(
	client *mongo.Client,
	collection *mongo.Collection,
	twoFaSmsColl *mongo.Collection,
	userRoleCollection *mongo.Collection,
) *Repository {
	return &Repository{
		Client:             client,
		Collection:         collection,
		TwoFaSmsColl:       twoFaSmsColl,
		UserRoleCollection: userRoleCollection,
	}
}
