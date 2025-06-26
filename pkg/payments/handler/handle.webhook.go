package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"com.fernando/pkg/app/ecomm/payment/model"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/webhook"
)

// * checkout.session.completed o payment_intent.succeeded me parece que es para web y mobile ver
// ver como testearlo
// habilitar en stripe payment_intent.succeeded y payment_intent.payment_failed
// ! ya agregaste el pago sucess en la base de datos perofalta el id, y is es stripe apple o google
// provider_payment_id el id por cada pago
func (h Handler) HandleStripeWebhook(w http.ResponseWriter, r *http.Request) {
	log.Println("Webhook recibido")

	const MaxBodySize = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodySize)
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		//w.WriteHeader(http.StatusServiceUnavailable)
		http.Error(w, "Error reading body", http.StatusServiceUnavailable)
		return
	}

	stripeSS := os.Getenv("STRIPE_SIGNING_SECRET")
	if err != nil {
		// ver como manejar el error
	}
	// * esto seria el servicio?
	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"), stripeSS)
	if err != nil {
		// esta bien poner err.Err() en que casos por que ira a stripe
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
		//http.Error(w, "Invalid signature", http.StatusBadRequest)
		//return
	}
	//fmt.Printf("Event: %v\n", event)
	//fmt.Printf("Event type: %T\n", event.Type)
	//fmt.Printf("Event type: %v\n", event.Type)

	switch event.Type {
	// * Agregar esto a los permisos del webhook ver los demas a agregar
	// * me parece que hay webhooks par amobile y web(checkout...)
	case "payment_intent.succeeded":
		//fmt.Printf("Pase por aqui 1\n")
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		paymentID := paymentIntent.Metadata["payment_id"]
		if paymentID == "" {
			//fmt.Printf("payment_id null\n")
		}
		//fmt.Printf("--------------------")
		//fmt.Printf("Payment ID: %s\n", paymentID)
		//fmt.Printf("--------------------")
		err = h.PaymentSrv.UpdateStatus(context.Background(), paymentID, model.PaymentPaid)
		if err != nil {
			//fmt.Printf("Err %v\n", err)
		}
		//fmt.Printf("Payment Intent Success: %v\n", paymentIntent)
		// manejar pago exitoso (ej. actualizar ba se de datos)
	case "payment_intent.payment_failed":
		//fmt.Printf("Pase por aqui 2\n")

		// manejar pago fallido grafana?
		// informar al usurio?

		// ! me parece que es para web
	case "checkout.session.completed":
		//fmt.Printf("Pase por aqui 3\n")
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			log.Printf("Error parsing session object: %v\n", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		//orderID := session.ClientReferenceID // O puedes usar session.Metadata["order_id"]
		//fmt.Printf("Session %v\n", session)
		//fmt.Printf("OrderID %v\n", orderID)

		// Aquí actualizas tu orden a pagado
		/* err := marcarOrdenComoPagada(orderID, session.ID)
		   if err != nil {
		       log.Printf("No se pudo actualizar la orden: %v\n", err)
		       http.Error(w, "Database error", http.StatusInternalServerError)
		       return
		   } */

		log.Println("✅ Orden actualizada correctamente")
	}

	w.WriteHeader(http.StatusOK)
}
