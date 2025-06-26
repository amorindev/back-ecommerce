package service

import (
	"com.fernando/pkg/app/stripe_customer/port"
	pProviderPort "com.fernando/pkg/payments/port"
)

var _ pProviderPort.PaymentProviderSrv = &Service{}

type Service struct {
	StripeCustomerRepo port.StripeCustomerRepo
}

func NewService(stripeCustomerRepo port.StripeCustomerRepo) *Service{
	return &Service{
		StripeCustomerRepo: stripeCustomerRepo,
	}
}
