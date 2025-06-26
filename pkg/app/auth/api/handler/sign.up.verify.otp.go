package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/pkg/app/auth/api/core"
	"com.fernando/pkg/app/auth/validate"
)

func (h Handler) SignUpVerifyOtp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req core.VerifyEmailOTP

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	defer r.Body.Close()

	_, err = validate.ValdiateVerifyOtp(req.OtpID, req.OtpCode, req.UserID, req.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	user,session, err := h.AuthService.SignUpVerifyOTP(context.Background(), req.OtpID, req.OtpCode, req.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	resp := core.AuthResp{
		User: user,
		Session: session,
	}

	// ? que se debería retornar todo? o lo que se modificó? de momento solo el user
	// * Retornar la session
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
