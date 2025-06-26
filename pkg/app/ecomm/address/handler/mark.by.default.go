package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
)

func (h Handler) MarkByDefault(w http.ResponseWriter, r *http.Request) {
	// * ver que retorna y que recibe
	// * me parece que no recibe el id del address por defecto si no lo buscamos desde
	// * la base de datos y que retorne los dos address el actualizado y el creado
	// * igualmente en (para redibujar la pantalla  en flutter)
	// *  address en CreateHandler MarkByDefaultaHandler
	// *  phones en CreateHandler MarkByDefaultaHandler
	
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	/* id, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.ErrorMessage{Msg: "Invalid id"})
		return
	} */

	// ! ver todos los nadlers que se retorna envia los ids no van en el body
	// ! lo mismo para user ver los status los errores

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
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "is_default field is required"})
		return
	}

	defer r.Body.Close()

	// deberia ser puntero para validar si es nil?
	/* if req.AddressID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "address_id is required"})
		return
	} */
	err = h.AddressSrv.ChangeDefault(context.Background(), id, *req.IsDefault)
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
