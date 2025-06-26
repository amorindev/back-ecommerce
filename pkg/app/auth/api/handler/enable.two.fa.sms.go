package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	md "com.fernando/cmd/api/middlewares"
	"com.fernando/internal/claim"
)

// ! falta en el servicio verifcar que no se duplique el numero ver en cuales
// hay dos cosas cuando el usuario agrega un nuevo nueor a la app por ejemplo en facebook
// se envia un sms para verifcar si es del user?
// en la ui el usuario selecciona el phone sin importar si esta verifcado(depende de lo que se defia arriba)
// verificado por que el usuario seleccionara uno y siempre enviará un sms
// agregar al otp metadata para marcar como verificado el phone?
// como diseñar la base de datos para agregar las distiontos tipos de twofa
// si es con phone agregar la relaccion con phoneid
func (h Handler) EnableTwoFaSms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// de momento solo el asignar el crear me parece facebook lo hace desde otr pantalla
	// u otro handler
	type reqBody struct {
		PhoneID string `json:"phone_id"`
	}
	var req reqBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	defer r.Body.Close()

	accessTokenClaim, ok := r.Context().Value(md.AccessTokenClaimsIDKey).(*claim.AccessTokenClaims)
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
	// * validaciones

	// vamos a decir si el phone no ha sido verificado retornar el otp id
	// pasar el phone id
	otpID, err := h.AuthService.EnableTwoFaSms(context.Background(), accessTokenClaim.UserID, req.PhoneID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	type respBody struct {
		OtpID string `json:"otp_id"`
	}
	resp := respBody{
		OtpID: otpID,
	}

	w.WriteHeader(http.StatusOK)
	// ! ver la response
	json.NewEncoder(w).Encode(resp)
}
