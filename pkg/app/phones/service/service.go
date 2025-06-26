package service

import "com.fernando/pkg/app/phones/port"

var _ port.PhoneSrv = &Service{}

type Service struct {
	PhoneRepo port.PhoneRepo
	PhoneTx   port.PhoneTransaction
}

func NewService(phoneRepo port.PhoneRepo, phoneTx port.PhoneTransaction) *Service {
	return &Service{
		PhoneRepo: phoneRepo,
		PhoneTx:   phoneTx,
	}
}
