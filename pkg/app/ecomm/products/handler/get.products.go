package handler

import (
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/pkg/app/ecomm/products/model"
)

func (h Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	products, err := h.ProductSrv.Get(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	type productResp struct {
		Products []*model.Product `json:"products"`
	}

	resp := productResp{
		Products: products,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
