package service

import paymentPort "com.fernando/pkg/app/ecomm/payment/port"

var _ paymentPort.PaymentSrv = &Service{}

type Service struct{
	PaymentRepo paymentPort.PaymentRepository
}

func NewPaymentSrv(paymentRepo paymentPort.PaymentRepository) *Service{
	return &Service{
		PaymentRepo: paymentRepo,
	}
}