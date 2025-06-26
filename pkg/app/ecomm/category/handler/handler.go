package handler

import (
	"net/http"

	"com.fernando/pkg/app/ecomm/category/port"
)

type Handler struct {
	CategorySrv port.CategoryService
}

func NewHandler(server *http.ServeMux, categorySrv port.CategoryService) *Handler{
	h := &Handler{
		CategorySrv: categorySrv,
	}

	server.HandleFunc("GET /v1/categories", h.GetCategories)
	return h
}
