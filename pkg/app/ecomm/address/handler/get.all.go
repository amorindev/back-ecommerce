package handler

import (
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	md "com.fernando/cmd/api/middlewares"
	"com.fernando/internal/claim"
	"com.fernando/pkg/app/ecomm/address/model"
)

// ! me parece innecesario retornar le useID dentro de la entidad al responder ver que otras entidades usan la refrerencia de otra tabla como order phone que se considera innecesario
func (h Handler) GetAll(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
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


	addresses, err := h.AddressSrv.GetAll(r.Context(),accessTokenClaim.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	type respBody struct{
		Addresses []*model.Address `json:"addresses"`
	}

	// no se si deve ser puntero
	resp := respBody{
		Addresses: addresses,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

