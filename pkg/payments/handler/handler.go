package handler

import (
	"net/http"

	"com.fernando/cmd/api/middlewares"
	"com.fernando/pkg/app/ecomm/payment/port"
	stripePort "com.fernando/pkg/payments/port"
)

type Handler struct {
	PaymentStripeSrv stripePort.PaymentProviderSrv
	PaymentSrv       port.PaymentSrv
}

func NewHandler(server *http.ServeMux, paymentStripeSrv stripePort.PaymentProviderSrv, paymentSrv port.PaymentSrv) *Handler {
	h := &Handler{
		PaymentStripeSrv: paymentStripeSrv,
		PaymentSrv:       paymentSrv,
	}

	// * o poner en el order que pasa con minio cuando dejarlo como servicio aprate
	// * y cuando incluirlo al DDD, en el caso del email puede ir separado y junto
	// * por ejemplo desde DDD puede necesitar order para decir hey hiciste una compra
	// * y para envviar email de confirmacion de correo, pero tambien sabemos que la logica
	// * es similar por eso de momento esta separado, ver que Servicios como payment file-storage
	// * necesitan capa servicio o solo adapter lo que pasa tambien es que los pagos en paypal
	// * y stripe pueden variar en handler y lo demas por eso no se agrega a order DDD
	pHandler := middlewares.LoggerMiddleware(middlewares.AuthMiddleware(h.CreatePaymentIntent))
	p2Handler := middlewares.LoggerMiddleware(middlewares.AuthMiddleware(h.HandlePaymentSheet))
	webhookH := middlewares.LoggerMiddleware(h.HandleStripeWebhook)
	server.HandleFunc("POST /v1/payments/create-payment-intent", pHandler)
	server.HandleFunc("POST /v1/payments/create-payment-intent2", p2Handler)
	// que reglas y validaciones va tener para que  solo stripe acceda a este handler
	// el weebhook funcionaría para ambos mobile web?
	server.HandleFunc("POST /v1/payments/stripe/webhook", webhookH)
	return h
}
