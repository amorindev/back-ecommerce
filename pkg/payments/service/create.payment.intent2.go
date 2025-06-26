package service

import "github.com/stripe/stripe-go/v82"

func (s *Service) CreatePaymentIntent2(amount int, currency string, metadata map[string]string, email string) (*stripe.PaymentIntent, error) {
	return nil,nil
}
