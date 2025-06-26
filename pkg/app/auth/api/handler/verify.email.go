package handler

import (
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/cmd/api/middlewares"
)

// en este y los demás handler solo enviar el token?

// saber que verifica el mmiddlewareauth
func (h Handler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	// deberria tener todo el token o solo el id del usuario
	userID, ok := r.Context().Value(middlewares.UserIDKey).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "userID invaid type"})
		return
	}

	// deberia validar el userID

	defer r.Body.Close()

	// verifica rporpósito del token

	// add logguers

	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "userID not found in the context"})
		return
	}

	// actuaizar el estado del uaurio email_verified = true
	// obetener el id del token
	err := h.AuthService.ConfirmEmail(userID)
	if err != nil {
		// deberia registrase el error, que errores se deben retornar?, se retorna al frontend?
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	// !
	// * Retornar la session igual en verify.otp.go
	// !
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(nil)
}
