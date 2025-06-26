package handler

import (
	"net/http"

	md "com.fernando/cmd/api/middlewares"
	"com.fernando/pkg/app/phones/port"
)

// ! asignar user id del token al crear
type Handler struct {
	PhoneSrv port.PhoneSrv
}
// por que esta entidad no puede ir dentro de handler users 
func NewHandler(server *http.ServeMux, phoneSrv port.PhoneSrv) *Handler {
	h := &Handler{
		PhoneSrv: phoneSrv,
	}
	// ! ver que middleware se ejecuta primero
	getAllH := md.LoggerMiddleware(md.AuthMiddleware(h.GetAll))

	createH := md.LoggerMiddleware(md.AuthMiddleware(h.Create))

	markByDefaultH := md.LoggerMiddleware(md.AuthMiddleware(h.MarkByDefault))

	// * no seria mejor la ruta user/phones
	// * o user/phoneid/phone ver

	server.HandleFunc("GET /v1/phones", getAllH)
	server.HandleFunc("POST /v1/phones", createH)
	server.HandleFunc("PATCH /v1/phones/{id}/default", markByDefaultH)
	return h
}
