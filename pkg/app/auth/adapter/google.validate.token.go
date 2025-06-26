package adapter

import (
	"context"
	"fmt"
)

// es tokenID
// Responsable de validar el ID tokenproporcionado por Google y extraer la información del usuario
func (a *Adapter) GoogleValidateToken(ctx context.Context, token string) (string, error) {
	// esto no debe estar dentro de GoogleProvider o service?
	idToken, err := a.GoogleVerifier.Verify(context.Background(), token)
	if err != nil {
		// invalid google token errors del user? o internal o unauthorized?
		return "", fmt.Errorf("auth adapter: GoogleValidateToken err %w", err)
	}


	// Extraer claims del token
	var claims struct {
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		//Name          string `json:"name"`
		//Picture       string `json:"picture"`
	}

	// ? es una interaccón con la api retorna estos valores
	if err := idToken.Claims(&claims); err != nil {
		// cannot parse token claims
		return "", fmt.Errorf("auth adapter - GoogleValidateToken: %w", err)
	}

	// validar email
	if !claims.EmailVerified {
		return "", fmt.Errorf("auth adapter - GoogleValidateToken, email not verified")
	}

	// si es picure con names return &Claim, nil
	return claims.Email, nil
}

/* func Test() {
	goUser, err := gothic.CompleteUserAuth(res,req)

	goth.UseProviders(
		google.New(cl)
	)
	// ! Accidentalmetne modifique el paqiete goth  ver nuevamente descargar porseacaso
}
*/
