package handler

import (
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	md "com.fernando/cmd/api/middlewares"
	"com.fernando/internal/claim"
	"com.fernando/pkg/app/ecomm/orders/model"
)

func (h Handler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
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

	orders, err := h.OrderSrv.GetAll(r.Context(), accessTokenClaim.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}
	type respBody struct {
		Orders []*model.Order `json:"orders"`
	}

	// ver como hacer el de lso otros get all
	var resp respBody
	// *estandar si no hay [] nulo es cargando
	// * o puede ser nulo pero en la ui adignnarle []
	// ! cambiar aqui y en los otros nil es un errror 
	// ! si no hay elementos []*model.Order{} ver
	if orders == nil {
		resp = respBody{
			Orders: []*model.Order{},
		}
	} else {
		// no se si deve ser puntero &respBody
		resp = respBody{
			Orders: orders,
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

