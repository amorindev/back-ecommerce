package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/cmd/api/middlewares"
	"com.fernando/internal/claim"
)

// obtiene el usuario
func (h Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
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

	// * mejor solo en dejarlo en UserHandler aqui estamos pasando el userID ver como gestionarlo
	// * ver si al hacer sign in retornaremos el provider sino seria solo GetUser
	// * no me parece de momento necesario el provider con el que inicio session 
	user, err := h.UserSrv.GetUser(context.Background(), accessTokenClaim.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	// ! recuerda que no se muestre los accesss rerfeshtokens y otp_id
	json.NewEncoder(w).Encode(user)

}
