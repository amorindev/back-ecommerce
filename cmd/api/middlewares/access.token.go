package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"com.fernando/cmd/api/message"
	"com.fernando/internal/claim"
)

// * un tipo por cada key
type userIDContext string
type tokenTypeContext string
type accessTokenClaimContext string

const UserIDKey userIDContext = "userID"
const TokenTypeIDKey tokenTypeContext = "token-type"
const AccessTokenClaimsIDKey accessTokenClaimContext = "access-token"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	accessString := os.Getenv("JWT_SIGNING_STRING")
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenString, err := tokenFromAuthorization(authHeader)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
			return
		}

		c, err := claim.GetAccessTokenFromJWT(tokenString, accessString)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, c.Subject)
		ctx = context.WithValue(ctx, TokenTypeIDKey, c.TokenType)
		ctx = context.WithValue(ctx, AccessTokenClaimsIDKey, c)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func tokenFromAuthorization(authorization string) (string, error) {
	if authorization == "" {
		// http.StatusUnauthorized
		return "", errors.New("authorization header is required")
	}

	if !strings.HasPrefix(authorization, "Bearer") {
		return "", errors.New("invalid authorization format")
	}

	l := strings.Split(authorization, " ")
	if len(l) != 2 {
		return "", errors.New("invalid authorization format")
	}

	return l[1], nil
}
