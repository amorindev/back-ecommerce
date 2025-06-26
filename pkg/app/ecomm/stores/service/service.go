package service

import "com.fernando/pkg/app/ecomm/stores/port"

var _ port.StoreSrv = &Service{}

type Service struct {
	StoreRepo port.StoreRepo
	StoreTx   port.StoreTx
}

func NewService(storeRepo port.StoreRepo, storeTx port.StoreTx) *Service {
	return &Service{
		StoreRepo: storeRepo,
		StoreTx: storeTx,
	}
}
