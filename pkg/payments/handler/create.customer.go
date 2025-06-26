package handler

import (
	//"encoding/json"
	"net/http"

	/* "com.fernando/cmd/api/message"
	"com.fernando/pkg/app/auth/api/core" */
)

// * El user_id y el email desde el token
type CreateCustomerReq struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}

type CreateCustomerResp struct {
	CustomerID string `json:"customer_id"`
}

func (h Handler) CreateStripeCustomer(w http.ResponseWriter, r *http.Request) {
	/* w.Header().Set("Content-Type", "application/json")
	var req CreateCustomerReq

	var u core.SignUpRequest

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	defer r.Body.Close()

	// * validar que  los campos no sean nulos, ver por que si se que en el token
	// * los estoy pasando no me deberia dar errror

	// 1. Validar si ya existe un customer asociado a este usuario
	existingCustomer, err := h.C */
}
