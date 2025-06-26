package handler

import (
	"net/http"

	"com.fernando/pkg/app/ecomm/products/port"
)

type Handler struct {
	ProductSrv port.ProductSrv
}

func NewHandler(server *http.ServeMux, productSrv port.ProductSrv) *Handler {
	h := &Handler{
		ProductSrv: productSrv,
	}
	// * Incluye variciones y sin variaciones
	server.HandleFunc("GET /v1/products", h.GetAllProducts)
	server.HandleFunc("GET /v1/products2", h.GetProducts)
	// * De momento separaremos la creacion con variantes y sin
	server.HandleFunc("POST /v1/products", h.Create)
	server.HandleFunc("POST /v1/product-variants", h.CreateWithVariations)

	return h
}
