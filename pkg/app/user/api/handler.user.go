package handler

import (
	"net/http"

	"com.fernando/cmd/api/middlewares"
	"com.fernando/pkg/app/user/port"
)

type Handler struct {
	UserSrv port.UserSrv
}

func NewHandler(server *http.ServeMux, userSrv port.UserSrv) *Handler {
	h := &Handler{
		UserSrv: userSrv,
	}

	userH := middlewares.LoggerMiddleware(middlewares.AuthMiddleware(h.UserHandler))
	server.HandleFunc("GET /v1/users/me", userH)

	return h
}