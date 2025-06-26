package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"com.fernando/cmd/api/message"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/customer"
	"github.com/stripe/stripe-go/v82/ephemeralkey"
	"github.com/stripe/stripe-go/v82/paymentintent"
)

// * payment intent para guardar datos de la tarjeta
// * 1. Crear el payment intent
// * 2. Genere un Customer, si aún no existe (deberia en ahand para al iniciar session le apresca la pantalla)
// * 3. Cree un ephemeral key
// * 4. Devuelva los tres valores: paymentIntent, customer, ephemeralKey
//
// ! que pasa si exsste el cliente ver mejor es separado como verificar que existe
func (h Handler) HandlePaymentSheet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stripeSK := os.Getenv("STRIPE_SECRET_KEY")
	if stripeSK == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "STRIPE_SECRET_KEY environment variable is not set"})
		return
	}
	stripe.Key = stripeSK

	// Crear un cliente de stripe
	// ver chatgtp save 5-29
	email := "fernan@gmail.com"
	params := &stripe.CustomerParams{
		Email: &email,
	}
	cust, err := customer.New(params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	// Crear una clave efímera para el cliente
	ephemeralKeyParams := &stripe.EphemeralKeyParams{
		Customer: stripe.String(cust.ID),
	}

	ephemeralKeyParams.StripeVersion = stripe.String("2025-03-31.basil")
	ek, err := ephemeralkey.New(ephemeralKeyParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}
	//fmt.Printf("Ok: %v\n", ek)

	// * Create payment intent
	amount := 562
	currency := "USD"
	piParams := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: &currency,
		Customer: stripe.String(cust.ID),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	pi, err := paymentintent.New(piParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	// responder con los datos necesarios
	resp := struct {
		PaymentIntent string
		EphemeralKey  string
		Customer      string
	}{
		PaymentIntent: pi.ClientSecret,
		EphemeralKey:  ek.Secret,
		Customer:      cust.ID,
	}
	json.NewEncoder(w).Encode(resp)
}
