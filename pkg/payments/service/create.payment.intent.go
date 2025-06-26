package service

import (
	"errors"
	"os"

	"github.com/stripe/stripe-go/v82"
	//"github.com/stripe/stripe-go/v82/customer"
	"github.com/stripe/stripe-go/v82/paymentintent"
)

// ! currency debe ser tipo,  currency stripe.Currency
func (s *Service) CreatePaymentIntent(amount int, currency string, metadata map[string]string) (*stripe.PaymentIntent, error) {
	stripeSK := os.Getenv("STRIPE_SECRET_KEY")
	if stripeSK == "" {
		return nil, errors.New("STRIPE_SECRET_KEY environment variable is not set")
	}
	stripe.Key = stripeSK

	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(int64(amount)),

		Currency: stripe.String(string(currency)),
		// AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
		// 	Enabled: stripe.Bool(true),
		// },

		// * Causa error ver
		/* TransferData:  &stripe.PaymentIntentTransferDataParams{
			Destination: stripe.String(""),
		}, */
		Metadata: metadata,
	}
	return paymentintent.New(params)
}

// ? cual es el flujo de save card con el customer en mobile y web y que datos se guardan en la
// ? base de datos
// * en stripe checkout para web ? y ademas que aprasca por defecto peru

/* func stripeWeb() {
	// Si estás usando Stripe Checkout con session, el país depende del campo customer o customer_email. Stripe intenta prellenar datos si el cliente ya existe y tiene país asignado.

	//🔧 Solución: Asocia un cliente con país Perú

   //En tu backend de Go al crear la sesión:

   params := &stripe.CheckoutSessionParams{
		Customer: stripe.String("cus_xxx"), // ID del cliente que ya tiene el país Perú
		// o si no tiene uno
		CustomerEmail: stripe.String("@gmail.com"),
	}

	//Para que funcione bien:
	//Asegúrate de que el cliente de Stripe tenga el campo address.country = "PE".
    //Opcionalmente, puedes crear el cliente así:
	// * de momento lo dejo aqui
	customerParams := &stripe.CustomerParams{
		Email: stripe.String("@gmail.com"),
		Address: &stripe.AddressParams{
			Country: stripe.String("PE"),
		},
	}
	customer, err := customer.New(customerParams)
	//Y luego usar ese customer.ID al crear la sesión de Checkout.


}

// ✅ 2. Para Stripe Elements (Payment Element en Flutter o webview)

func stripeMobileElements(){} */

// es payment intent (mobile) checkout session (web)
// y retorna un client secret
