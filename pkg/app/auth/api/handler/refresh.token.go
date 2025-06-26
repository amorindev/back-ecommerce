package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/cmd/api/middlewares"
	"com.fernando/internal/claim"
	"com.fernando/pkg/app/auth/api/core"
)

// * create token cuando se usa? - login - en que serrvicios

func (h Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	refreshTokenClaims, ok := r.Context().Value(middlewares.RefreshTokenClaimsKey).(*claim.RefreshTokenClaims)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "failed to parse claims - RefreshTokenHandler"})
		return
	}

	// * Validar
	// en este punto la purpose es correcta por que estamos usando un midleware una direrente firma
	// solo para el refreshtoken JWT_REFRESH_STRING ver asi para los demas
	if refreshTokenClaims.TokenType != "refresh-token" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "invalid-purpose RefreshToken"})
		return
	}

	accessToken, refreshToken, err := h.SessionService.RefreshToken(context.Background(), refreshTokenClaims.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(core.RefreshTokenResp{AccessToken: accessToken, RefreshToken: refreshToken})
}

// cuando cierra sesion eliminar e token, revocar
//refresh token tiene un midleware? o autorización

// verificar el refresh token con el JWT_REFRESH_TOKEN env
