package service

import "com.fernando/pkg/app/ecomm/address/port"

var _ port.AddressSrv = &Service{}

type Service struct {
	AddressRepo port.AddressRepo
	AddressTx port.AddressTx
}

func NewService(addressRepo port.AddressRepo,addressTx port.AddressTx) *Service {
	return &Service{
		AddressRepo: addressRepo,
		AddressTx: addressTx,
	}
}
