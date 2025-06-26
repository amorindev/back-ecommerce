package mongo

import (
	pConfigPort "com.fernando/pkg/app/ecomm/product-config/port"
	productItemPort "com.fernando/pkg/app/ecomm/product-item/port"
	productPort "com.fernando/pkg/app/ecomm/products/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ productPort.ProductTransaction = &Transaction{}

type Transaction struct {
	Client            *mongo.Client
	ProductRepo       productPort.ProductRepo
	ProductItemRepo   productItemPort.ProductItemRepo
	ProductConfigRepo pConfigPort.ProductConfigRepo
}

func NewTransaction(
	client *mongo.Client,
	productRepo productPort.ProductRepo,
	productItemRepo productItemPort.ProductItemRepo,
	pConfigRepo pConfigPort.ProductConfigRepo,
) *Transaction {
	return &Transaction{
		Client:            client,
		ProductRepo:       productRepo,
		ProductItemRepo:   productItemRepo,
		ProductConfigRepo: pConfigRepo,
	}
}
