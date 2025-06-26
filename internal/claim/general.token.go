package claim

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ! para cassos como forgot password casos criticos usar OTP, por que purpose es propenso a errores
// ! ver el apunte de deepsek en notion

// GeneralTokenClaims - Para casos donde necesites ambos (aunque generalmente es mejor usar las específicas)
type GeneralTokenClaims struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email,omitempty"`
	TokenType string `json:"token_type,omitempty"`
	jwt.RegisteredClaims
}

// ? crear funciones separadas o unirlas
// ! ver que se debe pasar por parámetro por que es de uso general
func NewGeneralToken(userID string, email string, purpose string, timeDuration time.Duration) *GeneralTokenClaims {
	return nil
}

// * Params
// * - purpose "email-verification"

// * Requisitos 1H maximo
func NewEmailVerificationToken(userID string, email string) *GeneralTokenClaims {
	return &GeneralTokenClaims{
		UserID:    userID,
		Email:     email,
		TokenType: "email-verification",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
}

// el otro es de reset cuando el usuario sabe su contraseña anterior
func NewForgotPasswordToken(userID string, email string) *GeneralTokenClaims {
	return &GeneralTokenClaims{
		UserID:    userID,
		Email:     email,
		TokenType: "forgot-password",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
}


func (c *GeneralTokenClaims) GetToken(accessString string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(accessString)) // o manejar el error y devover error personalizado
}

// usar el mismo accestoken
func GetGeneralTokenFromJWT(tokenString, accessString string) (*GeneralTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &GeneralTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(accessString), nil
	})

	if err != nil {
		return nil, fmt.Errorf("parse token error: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("token valid: false")
	}

	claim, ok := token.Claims.(*GeneralTokenClaims)
	if !ok {
		return nil, errors.New("invalid claim")
	}

	// ? que vaidaciones adicionales y que no van encada Get...FromJWT
	if claim.Subject == "" {
		return nil, errors.New("user id not found")
	}

	return claim, nil
}
