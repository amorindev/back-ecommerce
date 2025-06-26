package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"go.uber.org/zap"
	//"google.golang.org/api/idtoken"
)

// * se enviará nuestor propio token para las demás request y el client public del proveedor para su uso
// * en el frontend
// necesita su propio midlleware?
// como afecta a auth middleware
// login/register
func (h Handler) GoogleSignIn(w http.ResponseWriter, r *http.Request) {
	logger := zap.Must(zap.NewProduction())
	w.Header().Set("Content-Type", "application/json")

	var bodyReq struct {
		RememberMe    bool   `json:"remember_me"`
		GoogleTokenID string `json:"token_id"` // o desde el header?
	}

	defer logger.Sync()

	if err := json.NewDecoder(r.Body).Decode(&bodyReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	//red := "\033[31m"
	reset := "\033[0m"
	yellow := "\033[33m"
	//green := "\033[32m"

	logger.Info("--------------- idToken ----------------")
	logger.Info(yellow + bodyReq.GoogleTokenID + reset)
	logger.Info(yellow + bodyReq.GoogleTokenID + reset)
	logger.Info("----------------------------------------")
	
	// * funcion para validar los campos o  crear GoogleSignIn constructor y retornar el error
	// * Cuál es mejor?
	if bodyReq.GoogleTokenID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "google token_id is required"})
		return
	}

	auth, err := h.AuthService.GoogleSignIn(context.Background(), bodyReq.RememberMe, bodyReq.GoogleTokenID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(auth)
}
