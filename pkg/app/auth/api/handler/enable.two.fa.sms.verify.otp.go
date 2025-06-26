package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/cmd/api/middlewares"
	"com.fernando/internal/claim"
)

func (h Handler) EnableTwoFaSmsVerifyOtp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accessTokenClaim, ok := r.Context().Value(middlewares.AccessTokenClaimsIDKey).(*claim.AccessTokenClaims)
	if !ok {
		// bad request o unhotorixed o internal server?
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "failed to parse claims - AccessTokenClaims"})
		return
	}

	if accessTokenClaim.Subject == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "user-not-found-on-claim"})
		return
	}

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

	// * validar la request
	if req.OtpCode == "" || req.OtpID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "otp_code or otp_id are requireds"})
		return
	}

	// TODO: tenemos varias cosas una enos recomendable creo retornar el userTwoFaSmsID
	// TODO: y nuevamente enviarlo, ver cuales la relaccion UserTwoFaSmsID, me parece
	// TODO: que es 1 a 1 pero mongo siempre ponbra un id por defecto, entnces
	// TODO:consideraciones socbre la collecion de mongo esta en el readmen de auth
	// TODO: dice userTwoFaSms.ID = userID el mismo id deve si tener solo uno y docs
	// TODO: y la otra es agregar metadata en el otp
	// TODO: de momento warning dejaremo todo normal y pasaremos userid
	user,err := h.AuthService.EnableTwoFaSmsVerifyOtp(context.Background(), req.OtpID, req.OtpCode, accessTokenClaim.UserID)
	if err != nil {
	  	w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}
	// ! retornar el AuthResp con la session actualizada
	// ! por que el user se crea a partir de la session
	// ! siguiendo esta logica signupverifyotp tendria que retonar el Auth resp
	// ! ver por que tambien se debe retornar en el signup la session 
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
// ! ver esto por que antes de habilitar 
// ! que pasa si los reles de un usuario cambian
// ! ver que no afecte al signend state flutter

/*
🧩 ¿Qué pasa con Session?
🔐 Retorna una nueva session solo si:
Generaste un nuevo accessToken o refreshToken en la operación.

Tu backend invalida la sesión anterior o la considera obsoleta después del cambio.

❌ No retornes una nueva Session si el token no cambió.
Esto es importante porque si tu token sigue siendo válido y no se renueva, puedes seguir usando el mismo Session actual.
*/