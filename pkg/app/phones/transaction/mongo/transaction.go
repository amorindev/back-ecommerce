package mongo

import (
	"com.fernando/pkg/app/phones/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.PhoneTransaction = &Transaction{}

type Transaction struct {
	Client    *mongo.Client
	PhoneRepo port.PhoneRepo
}

func NewTransaction(client *mongo.Client, phoneRepo port.PhoneRepo) *Transaction {
	return &Transaction{
		Client:    client,
		PhoneRepo: phoneRepo,
	}
}
