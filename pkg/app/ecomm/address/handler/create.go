package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/cmd/api/middlewares"
	"com.fernando/internal/claim"
	"com.fernando/pkg/app/ecomm/address/core"
	"com.fernando/pkg/app/ecomm/address/model"
)

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
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
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "userid-not-found-on-claim"})
		return
	}

	var req core.CreateAddressReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	defer r.Body.Close()

	// asi como en flutter usamos ! en go seria *user.Label por que sabemos que debe venir
	// ver
	// ver cuales serán punteros
	if req.PostalCode == "" || req.Label == nil || req.AddressLine == "" || req.Country == "" {
		w.WriteHeader(http.StatusBadRequest)
		//json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Country,Label, PostalCode or Addreess line are required"})
		return
	}

	address := &model.Address{
		UserID:      accessTokenClaim.Subject,
		Label:       req.Label,
		AddressLine: req.AddressLine,
		State:       req.State,
		Country:     req.Country,
		City:        req.City,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		PostalCode:  req.PostalCode,
	}

	// de momento solo en get context de la request o no se si en todos y cuando agrupar ctx
	err = h.AddressSrv.Create(context.Background(), address)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)
}
