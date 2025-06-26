package adapter

import (
	"context"
	"errors"
)

func (a *Adapter) AppleValidateToken(ctx context.Context, token string) (string, error) {
	// Obtener claves públicas de apple
	/*keySet := oauth2.Endpoint{AuthURL: a.AppleProvider.AppleKeyUrl}

	// Decodificar el token
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("auth adapter - AppleValidateToken, unexpected signing method")
		}
		return keySet, nil
	})
	if err != nil {
		return "", err
	}
	var email string
	var name string
	var emailVerified bool
	//claims := &UserClaims{}
	if claimsMap, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
		email = claimsMap["email"].(string) // aqui no hace la validación
		name = claimsMap["name"].(string)
		emailVerified = claimsMap["email_verified"].(bool)
		//return claims, nil
	} */
	/*
		if claimsMap, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			claims.Email = claimsMap["email"].(string)
			claims.Name = claimsMap["name"].(string)
			claims.EmailVerified = claimsMap["email_verified"].(bool)
			return claims, nil
		}
	*/
	//fmt.Printf("Name %s\n", name)
	//fmt.Printf("Email verified %v\n", emailVerified)
	//return email, nil
	return "", errors.ErrUnsupported
}
