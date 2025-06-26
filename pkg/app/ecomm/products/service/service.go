package service

import (
	productPort "com.fernando/pkg/app/ecomm/products/port"
	fileStoragePort "com.fernando/pkg/file-storage/port"
)

var _ productPort.ProductSrv = &Service{}

type Service struct {
	ProductRepo    productPort.ProductRepo
	FileStorageSrv fileStoragePort.FileStorageSrv
	ProductTx      productPort.ProductTransaction
}

func NewService(
	productRepo productPort.ProductRepo, 
	fileStorageSrv fileStoragePort.FileStorageSrv, 
	productTx productPort.ProductTransaction,
) *Service {
	return &Service{
		ProductRepo:    productRepo,
		FileStorageSrv: fileStorageSrv,
		ProductTx:      productTx,
	}
}
