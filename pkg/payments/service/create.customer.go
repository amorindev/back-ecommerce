package service

import (
	"context"
	"time"

	"com.fernando/pkg/app/stripe_customer/errors"
	"com.fernando/pkg/app/stripe_customer/model"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/customer"
)

func (s *Service) CreateCustomer(ctx context.Context, userID string, email string) (string,error) {
	existingCustomer, err := s.StripeCustomerRepo.GetByUserID(ctx,userID)
	if err != nil && err != errors.ErrStripeCustomerNotFound {
	  return "", err
	}
	if existingCustomer != nil {
		return existingCustomer.CustomerID.(string), nil
	}

	// Crear un cutomer en stripe
	params := &stripe.CustomerParams{
		Email: stripe.String(email),
		//Name: ,
		// de momento no se para que
		Metadata: map[string]string{
			"user_id": userID,
		},
	}

	customer, err := customer.New(params)
	if err != nil {
	  return "", err
	}

	now:= time.Now()
	// Guardar en la base de datos
	// stripe customer pero me parece mejor 
	newSCustomer := &model.StripeCustomer{
		UserID: userID,
		CustomerID: customer.ID,
		CreatedAt: &now,
	}
	err = s.StripeCustomerRepo.Insert(ctx,newSCustomer)
	if err != nil {
	  return "", err
	}
	return customer.ID, nil
}