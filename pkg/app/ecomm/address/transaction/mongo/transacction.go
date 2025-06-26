package mongo

import (
	"com.fernando/pkg/app/ecomm/address/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.AddressTx = &Transaction{}

type Transaction struct{
	Client *mongo.Client
	AddressRepo port.AddressRepo
}

func NewTransaction(client *mongo.Client, addressRepo port.AddressRepo)*Transaction{
	return &Transaction{
		Client: client,
		AddressRepo: addressRepo,
	}
}