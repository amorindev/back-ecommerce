package port

import (
	"context"

	"github.com/stripe/stripe-go/v82"
)


type PaymentProviderSrv interface {
	//CreatePayment(amount float64, userID string, orderID string) (*stripe.CheckoutSession, error)
	//GetPaymentStatus(pID string) (*stripe.CheckoutSession, error)
	// ! cambiar a primiticos *stripe.Payment intent 
	CreatePaymentIntent(amount int, currency string, metadata map[string]string) (*stripe.PaymentIntent, error)
	//CreatePaymentIntent2(amount int, currency string, metadata map[string]string, email string) (*stripe.PaymentIntent, error)
	// * ver si se va agregar  el name tanto a user struct
	// * como a esta interface ver handler create.customer
	// * retorna el customerid
	CreateCustomer(ctx context.Context, userID string, email string) (string, error)
}
