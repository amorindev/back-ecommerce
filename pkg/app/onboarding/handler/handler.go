package handler

import (
	"net/http"

	"com.fernando/cmd/api/middlewares"
	"com.fernando/pkg/app/onboarding/port"
)

type Handler struct {
	OnboardingSrv port.OnboardingSrv
}

func NewHandler(server *http.ServeMux, onboardingSrv port.OnboardingSrv) *Handler {
	h := &Handler{
		OnboardingSrv: onboardingSrv,
	}
	// no implementado
	server.HandleFunc("GET /v1/onboardings", http.NotFound)
	server.HandleFunc("GET /v1/onboardings2", middlewares.LoggerMiddleware(h.GetOnboardings))
	return h
}
