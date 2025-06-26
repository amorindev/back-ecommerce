package handler

import (
	"net/http"

	md "com.fernando/cmd/api/middlewares"
	"com.fernando/pkg/app/ecomm/address/port"
)

type Handler struct {
	AddressSrv port.AddressSrv
}

func NewHandler(server *http.ServeMux, addressSrv port.AddressSrv) *Handler {
	h := &Handler{
		AddressSrv: addressSrv,
	}

	// no se hasta que punto es evidente GetAll o GetAllAddress Create o CreateAddress
	// lo mismo para los nombre de archivos y busqueda rápidaw
	getAllH := md.LoggerMiddleware(md.AuthMiddleware(h.GetAll))
	createH := md.LoggerMiddleware(md.AuthMiddleware(h.Create))
	markByDefaultH := md.LoggerMiddleware(md.AuthMiddleware(h.MarkByDefault))

	server.HandleFunc("GET /v1/addresses", getAllH)
	server.HandleFunc("POST /v1/addresses", createH)
	server.HandleFunc("PATCH /v1/addresses/{id}/default", markByDefaultH)
	return h
}
