package claim

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// RefreshTokenClaims - Claims para tokens de refresco
type RefreshTokenClaims struct {
	UserID    string `json:"user_id"`
	TokenType string `json:"token_type"` // Podemos especificar que es refresh, al acces también?
	jwt.RegisteredClaims
}


// ! Verificar que datos pondremos en el accesstoken y refreshtoken
// ! debemos buscar el refresh token mediante el refreshID
func NewRefreshToken(userID string, rememberMe bool) *RefreshTokenClaims {
	refreshID := uuid.New().String()
	expiresAt := time.Hour * 24 * 7
	if rememberMe {
		expiresAt = time.Hour * 24 * 30
	}
	return &RefreshTokenClaims{
		UserID:    userID,
		TokenType: "refresh-token",
		RegisteredClaims: jwt.RegisteredClaims{
			// como solo se envía el refresh token para actualizar los toknes
			// cómo eliminariamo el refresh token de la base de datos mejoar agregar el id
			// para con una go rutina eliminarlo, y tambien usar el cronjob
			// el dilema que teniamos pasar id, documentar estos  puntos
			// y el tipos de flujo de tokens larga y corta duración
			//Issuer: , ? poner o no
			// Audience: , ? poner o no
			ID:        refreshID,
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAt)),
		},
	}
}

func (c *RefreshTokenClaims) GetToken(refreshString string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(refreshString)) // o manejar el error y devover error personalizado
}

func GetRefreshTokenFromJWT(tokenString, refreshString string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(refreshString), nil
	})

	if err != nil {
		return nil, fmt.Errorf("parse token error: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("token valid: false")
	}

	claim, ok := token.Claims.(*RefreshTokenClaims)
	if !ok {
		return nil, errors.New("invalid claim")
	}

	// ? que vaidaciones adicionales y que no van encada Get...FromJWT
	if claim.Subject == "" {
		return nil, errors.New("user id not found")
	}

	return claim, nil
}
