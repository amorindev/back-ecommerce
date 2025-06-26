package handler

import (
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/pkg/app/auth/api/core"
)

// ! desde los constructores valores por defecto remeberme
// como mnejar los parámetros opcionales con * o ""
// si el remenber me es opcional validar?
func (h Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req core.SignInReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	defer r.Body.Close()

	// * validaciones
	err = req.IsSignInValid()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	user, session, otpID, err := h.AuthService.SignIn(r.Context(), req.Email, req.Password, req.RememberMe)
	if err != nil {
		// puede ser usernot found
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	resp := core.AuthResp{
		User:        user,
		Session:     session,
		Credentials: user.AuthProviderCreate,
		OtpID:       otpID,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
