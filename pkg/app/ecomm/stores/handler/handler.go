package handler

import (
	"net/http"

	md "com.fernando/cmd/api/middlewares"
	"com.fernando/pkg/app/ecomm/stores/port"
)

type Handler struct {
	StoreSrv port.StoreSrv
}

func NewHandler(server *http.ServeMux, storeSrv port.StoreSrv) *Handler {
	h := Handler{
		StoreSrv: storeSrv,
	}

	getAllH := md.LoggerMiddleware(md.AuthMiddleware(h.GetAll))

	server.HandleFunc("GET /v1/stores/", getAllH)

	return &h
}
