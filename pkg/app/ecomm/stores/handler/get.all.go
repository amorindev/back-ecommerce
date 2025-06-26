package handler

import (
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/pkg/app/ecomm/stores/model"
)

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stores, err := h.StoreSrv.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	// ! pasar solo el arreglo o dentro del objeto siempre asi
	type Resp struct {
		Stores []*model.Store `json:"stores"`
	}

	resp := Resp{
		Stores: stores,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
