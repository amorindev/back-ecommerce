package core

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// * Sign In
type SignInReq struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
	Platform   string `json:"platform"` // pra ver la lista de sesiones  como facebook
}

// ! crear un constructor para grpc los token se pasan por el context grpc

// * en este sentido si se puede usar validaciones mediante estructuras
func (req SignInReq) IsSignInValid() error {
	validate := validator.New()

	// * validar el email
	if req.Email == "" {
		return errors.New("email-field-is-required")
	}
	err := validate.Var(req.Email, "email")
	if err != nil {
		return fmt.Errorf("invalid email: %s", err.Error())
	}

	// Validar contraseña
	if req.Password == "" {
		return fmt.Errorf("password-field-is-required")
	}
    // ! Cual es la diferencia de valores opcionales para logica de negocio 
    // ! y valores opcionales para guardar en la base de datos
	return nil
}

type SignInResp struct {
	AccessToken  string          `json:"access_token"`
	RefreshToken string          `json:"refresh_token"`
	User         *SignInUserResp `json:"user"`
}

type SignInUserResp struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Username      string `json:"username"`
}
