package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
)

// ! el problema que tenemos aqui phone y address es que
// ! estamos modificando dos documentos o registros al que pasará aserpor defecto y el que ya no es por defecto entonces como actalizo mi ui
// ! al no retornar nada cambiar desde la ui los dos elemmentos

func (h Handler) MarkByDefault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	type reqBody struct {
		IsDefault *bool `json:"is_default"`
	}

	var req reqBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}
	if req.IsDefault == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "active field is required"})
		return
	}

	defer r.Body.Close()

	err = h.PhoneSrv.ChangeDefault(context.Background(), id, *req.IsDefault)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}
	// aqui sería estatus no content pero como manejarlo desde el frontend
	// de moment sensillo
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(status)  me dice chat que no se retorna nada unmessage? ver
}
