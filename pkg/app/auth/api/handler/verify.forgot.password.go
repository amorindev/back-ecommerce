package handler

import "net/http"

func (h Handler) VerifyForgotPassword(w http.ResponseWriter, r *http.Request){
	// obtener el claim y verificar el TokenType

	w.Header().Set("Content-Type","application/json")
}