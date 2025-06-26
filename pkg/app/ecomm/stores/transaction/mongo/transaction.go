package mongo

import (
	addressPort "com.fernando/pkg/app/ecomm/address/port"
	storePort "com.fernando/pkg/app/ecomm/stores/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ storePort.StoreTx = &Transaction{}

type Transaction struct {
	Client   *mongo.Client
	StoreRepo storePort.StoreRepo 
	AddressRepo addressPort.AddressRepo
}

func NewStoreTx(
	client *mongo.Client,
	storeRepo storePort.StoreRepo,
	addressRepo addressPort.AddressRepo,
) *Transaction {
	return &Transaction{
		Client:   client,
        StoreRepo: storeRepo,
        AddressRepo: addressRepo,
	}
}
