package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/cmd/api/middlewares"
	"com.fernando/internal/claim"
	"com.fernando/pkg/app/phones/core"
	"com.fernando/pkg/app/phones/model"
)

// ! tenemos el create si es 200 o no content desde el frontend debemos modificar el anterior y cambiarlo a que no es por defecto y agregar el creado para que tenga sentido
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

	var req core.CreatePhoneReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	defer r.Body.Close()

	if req.Number == nil || req.CountryCode == nil || req.CountryIsoCode == nil {
		w.WriteHeader(http.StatusBadRequest)
		//json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "number, countrycode countryIsocode are required"})
		return
	}

	phone := &model.Phone{
		UserID:         accessTokenClaim.Subject,
		Number:         req.Number,
		CountryCode:    req.CountryCode,
		CountryIsoCode: req.CountryIsoCode,
	}

	err = h.PhoneSrv.Create(context.Background(), phone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(phone)
}
