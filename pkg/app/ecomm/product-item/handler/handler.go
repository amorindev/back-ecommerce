package handler

import (
	"net/http"

	productItemPort "com.fernando/pkg/app/ecomm/product-item/port"
)

type Handler struct {
	ProductSrv productItemPort.ProductSrv
}

func NewHandler(server *http.ServeMux, productItemSrv productItemPort.ProductSrv) *Handler {
	h := &Handler{
		ProductSrv: productItemSrv,
	}

	// * Create handlers
	//server.HandleFunc("GET /v1/products", )
	//server.HandleFunc("POST /v1/products", )
	//server.HandleFunc("POST /v1/product-variations",)
	
	// *Add coment y marcar como favirito
	//server.HandleFunc("POST /v1/product-variations", h.CreateWithVariations)

	return h
}
