package service

import (
	productPort "com.fernando/pkg/app/ecomm/product-item/port"
	fileStgPort "com.fernando/pkg/file-storage/port"
)

var _ productPort.ProductSrv = &Service{}

// * adt adapter
type Service struct {
	ProductItemRepo    productPort.ProductItemRepo
	FileStorageSrv fileStgPort.FileStorageSrv
	// FIle storage name si se va a separar el storage
}

func NewService(productItemRepo productPort.ProductItemRepo, fileStgSrv fileStgPort.FileStorageSrv) *Service {
	return &Service{
		ProductItemRepo:    productItemRepo,
		FileStorageSrv: fileStgSrv,
	}
}
