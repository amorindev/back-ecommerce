package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// UserClaims representa los datos del usuario en Facebook
type FacebookUserClaims struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ValidateToken verifica el token de Facebook y extrae email y nombre
func (a *Adapter) FacebookValidateToken(ctx context.Context, token string) (string, error) {
	// Verificar token con Facebook
	url := fmt.Sprintf("https://graph.facebook.com/me?fields=id,name,email&access_token=%s", token)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("invalid facebook token")
	}

	var claims FacebookUserClaims
	if err = json.NewDecoder(resp.Body).Decode(&claims); err != nil {
		return "", errors.New("cannot parse token claims")
	}

	// validar email
	if claims.Email == "" {
		return "", errors.New("email not provided by facebook")
	}

	return claims.Email, nil
	//return &claims, nil
}
