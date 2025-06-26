package handler

import (
	"net/http"

	"com.fernando/pkg/app/ecomm/variation/port"
)

type Handler struct {
	VariationSrv port.VariationService
}

func NewHandler(server *http.ServeMux, variationSrv port.VariationService) *Handler {
	h := &Handler{
		VariationSrv: variationSrv,
	}

	server.HandleFunc("GET v1/variations/var-option", h.GetWithVarOptions)
	return h
}


