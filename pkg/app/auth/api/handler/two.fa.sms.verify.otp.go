package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/pkg/app/auth/api/core"
)

func (h Handler) TwoFaSmsVerifyOtp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type reqBody struct {
		OtpID   string `json:"otp_id"`
		OtpCode string `json:"otp_code"`
	}

    var req reqBody

    err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}
    defer r.Body.Close()

    if req.OtpCode == "" || req.OtpID == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(message.ErrorMessage{Message: "otpCode or otpID is required"})
        return
    }
    user, session, err := h.AuthService.TwoFaSmsVerifyOtp(context.Background(),req.OtpID, req.OtpCode)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}
    resp := core.AuthResp{
        User: user,
        Session: session,
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(resp)
}

