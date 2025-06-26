package handler

import (
	"net/http"
)

func (h Handler) SendEmailVerification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//claim.NewTokenEmailVerify(userID, time.Hour)
}
