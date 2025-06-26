package mongo

import (
	"com.fernando/pkg/app/ecomm/orders/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.OrderRepository = &Repository{}

type Repository struct {
	Client           *mongo.Client
	Collection       *mongo.Collection
	OrderProductColl *mongo.Collection
}

func NewRepository(client *mongo.Client, coll *mongo.Collection, orderProductColl *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: coll,
		OrderProductColl: orderProductColl,
	}
}
