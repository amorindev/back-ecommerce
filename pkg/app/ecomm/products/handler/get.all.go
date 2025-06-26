package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"com.fernando/cmd/api/message"
	"com.fernando/pkg/app/ecomm/products/model"
)

func (h Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// * en la no se ve refejado el page ver que pasa en la primera llamada
	// * aunque en el atio dara error o 0 y ahi ya lo manejamos
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	// * se puede separar pero es buena lógina si no eres tan extricto,
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	products, count, pages, err := h.ProductSrv.GetAll(r.Context(), limit, page)
	if err != nil {
		//http.Error() si esto respondemos el application/json al último
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	//reqParams := make(url.Values)

	type ProductResp struct {
		Count    int              `json:"count"`
		Pages    int              `json:"pages"`
		Next     *string          `json:"next"`
		Previus  *string          `json:"previus"`
		Products []*model.Product `json:"products"`
	}

	// ! se cacula el page y el count como? mas que nada el page
	// * revisar desde arriba tambien hay page y lo que retorna del  servicio cual de los dos usar
	// * y que conincida en capa de servicio tambien agregaste otras validaciones que funcione en conjunto
	// * ver la lógica de paginación
	var next *string
	nextStr := fmt.Sprintf("http://localhost:8000/v1/products?page=%d", page+limit)
	next = &nextStr
	if page == pages {
		next = nil
	}

	var previus *string
	previusStr := fmt.Sprintf("http://localhost:8000/v1/products?page=%d", page)
	previus = &previusStr
	// en la url no aparece page = 1 ver
	if page == 1 {
		previus = nil
	}

	// *Sacar del env- ver en ambos cuando ?page es igual a Pages next es null
	productResp := &ProductResp{
		Count: count,
		//Pages:

		Next:     next,
		Previus:  previus,
		Products: products,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productResp)
}

/*
_, err := a.MinioClient.PutObject(ctx, "ecomm-test", "products/"+fileName, fileReader, -1, options)
	if err != nil {
		return err
	}
*/
