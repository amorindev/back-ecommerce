package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"com.fernando/cmd/api/message"
	"com.fernando/internal/claim"
)

type refreshTokenClaimsContext string

const RefreshTokenClaimsKey refreshTokenClaimsContext = "refresh-token-claims"

func RefreshTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	refreshString := os.Getenv("JWT_REFRESH_STRING")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			RefreshToken string `json:"refresh_token"`
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if request.RefreshToken == "" {
			http.Error(w, "Refresh token required", http.StatusBadRequest)
			return
		}

		c, err := claim.GetRefreshTokenFromJWT(request.RefreshToken, refreshString)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
			return
		}
		ctx := context.WithValue(r.Context(), RefreshTokenClaimsKey, c)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
