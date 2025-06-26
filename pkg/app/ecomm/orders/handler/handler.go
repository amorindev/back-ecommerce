package handler

import (
	"net/http"

	"com.fernando/cmd/api/middlewares"
	"com.fernando/pkg/app/ecomm/orders/port"
)

type Handler struct {
	OrderSrv port.OrderService
}

func NewHandler(server *http.ServeMux, orderSrv port.OrderService) *Handler {
	h := &Handler{
		OrderSrv: orderSrv,
	}

	getAllH := middlewares.LoggerMiddleware(middlewares.AuthMiddleware(h.GetAllOrders))
	createH := middlewares.LoggerMiddleware(middlewares.AuthMiddleware(h.Create))

	// ver sistema de nombres variables en ingles
	server.HandleFunc("GET /v1/orders", getAllH)
	server.HandleFunc("POST /v1/orders", createH)
	return h
}
