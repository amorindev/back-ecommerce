package claim

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenType para que no se envien cualquier token con diferente propósito seguridad - add al refsh token
// refresh token tiene type
// AccessTokenClaims - Claims para tokens de acceso
type AccessTokenClaims struct {
	// cual se va usar UserID o lo adjuntamos en Subject ver para no tener dos
	// campos que realicen lo mismo
	UserID    string   `json:"user_id"` // user o auth - auth esta relacionado con la session
	Email     string   `json:"email"`
	Roles     []string `json:"role,omitempty"` // Campo opcional para roles
	TokenType string   `json:"token_type"`     // forgot-password - verify-account - sign-in
	jwt.RegisteredClaims
	// ! el id del auth pertenece a la seesion entonces tendrá varias sessiones varios tokens
	// el hash del token,
	// es el id de la seesion o auth
}

/*
jti	uuid.New().String()	ID único del token (para evitar replay attacks). falta agregar ver
*/
// ! Verificar que datos pondremos en el accesstoken y refreshtoken
// ? authID no seria en subjet para lasessiones o en el
func NewAccessToken(userID string, email string, issuer string, roles []string, audience []string) *AccessTokenClaims {
	
	return &AccessTokenClaims{
		UserID:    userID,
		Email:     email,
		Roles:     audience,
		TokenType: "sign-in",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,   // no se repitiría
			Issuer:    issuer,   // "https://localhost:5000",  dominio .com
			Audience:  audience, // []string{"https://api1.tudominio.com", "https://api2.tudominio.com"},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 2)),
		},
	}
}

func (c *AccessTokenClaims) GetToken(accessString string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(accessString))
}

// * no sería get claim
// signing por accesString
func GetAccessTokenFromJWT(tokenString, signingString string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingString), nil
	})

	if err != nil {
		if err.Error() == "token has invalid claims: token is expired" {
			return nil, errors.New("token-is-expired")
		}
		return nil, fmt.Errorf("parse token error: %w, type: %T", err, err)
	}

	if !token.Valid {
		return nil, errors.New("token valid: false")
	}

	claim, ok := token.Claims.(*AccessTokenClaims)
	if !ok {
		return nil, errors.New("invalid claim")
	}

	// ! no se si yo puse esta validacion pero esta bien el purpose si desde el handler por que no es
	// ! general

	if claim.Subject == "" {
		return nil, errors.New("user id not found")
	}

	return claim, nil
}
