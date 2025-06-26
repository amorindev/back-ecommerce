package core

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// * Sign Up
type SignUpReq struct {
	Email           string  `json:"email"`
	Name            *string `json:"name"`
	Username        *string `json:"username"`
	Phone           *string `json:"phone"`
	Password        string  `json:"password"`
	ConfirmPassword string  `json:"confirm_password"`
	WithCode        bool    `json:"with_code"`
}

// * Usamos validate funcion por que si hacemos en el UserCreate params no hay forma de reutilizarlo en gRPC
// * y no lo podemos agregar a Auth y usar un métod isSignInValid o asi por que sobrecargaríamos el estructura
// * (si hay solucion pasarlo como parámetro), por eso usamos isSignInvalid o signup
// * y algunos datos no se guardan en la base de datos o en el servicio solo son para verificar ConfirmPAssword,
// * y otris ver
// * revisar las otras validaciones trim validator 10
// ? aplicr trim a todos los que requieren
// * probar diferentes casos por ejemplo
// * que un campo es opcional string (primero validar que no se nil despes validar) boolen
// ? como realizar las validaciones
// ponemos * punteros en el core para qu no se guarde como "" y en la api se muestre null
// los campose requeridos si normal sin punteros
// * las validaciones como no debe existir dos usernames se realiza desde el servicio
// ? como validar agregados de core
// ! Como personalizar los errores
// ! y segun el  error que venga del paquite por ejemplo mongo ErrNotFound
// ! ver el de jwt el validator y asi errores personalizados
// ? se debe validar todo?
func (req SignUpReq) IsSignUpValid() error {
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email-field-is-required")
	}

	validate := validator.New()
	err := validate.Var(req.Email, "email")
	if err != nil {
		return fmt.Errorf("SignUpValidateErr: %v", err.Error())
	}
	if req.Name != nil {
		println(*req.Name)
		// validacion inenecesaria por que si viene "" go lo tranafirma en nil y no un puntero
		// a un campo ""
		//if strings.TrimSpace(*req.Name) == "" {}
		// esto tambien no seria util por que para que no sea nil debe tern almenos un caacterw
		if len(*req.Name) < 1 {
			return errors.New("name must be at least 1 characters")
		}
		// deberia ser no debe estar vacio
		// 
	}

	if req.Username != nil && strings.TrimSpace(*req.Username) == "" {
		return errors.New("username-field-is-required")
	}

	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password-field-is-required")
	}
	if strings.TrimSpace(req.ConfirmPassword) == "" {
		return errors.New("confirm-password-field-is-required")
	}
	// validar caracteres especiales y cantidad minima de letras
	if len(req.Password) < 8 {
		return errors.New("week-password")
	}
	// ? ver si es nulo flujo
	if req.Password != req.ConfirmPassword {
		return errors.New("passwords-do-not-match")
	}

	err = validate.VarWithValue(req.Password, req.ConfirmPassword, "eqfield")
	if err != nil {
		return errors.New("passwords-do-not-match")
	}
	/* if u.AuthProviderCreate != nil {
	        return  errors.New("add user data \"provider_data\" data ")
		} */
	// cantidad de catacteres desde punto env
	// me parece que no es necesario trim por que estamo evluando el len() solo al password y lo comp
	// comparamos con confirmpassword sin hacer nada
	return nil
}
