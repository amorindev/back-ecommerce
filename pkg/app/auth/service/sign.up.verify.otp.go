package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"com.fernando/pkg/app/auth/service/constants"
	sessionM "com.fernando/pkg/app/session/model"
	"com.fernando/pkg/app/user/model"
)

// ! cambiar por que el veridy otp puede ser para varios no solo para auth ver
// * Me parece que no necesitaremos userID
func (s *Service) SignUpVerifyOTP(ctx context.Context, otpID string, otpCode string, email string) (*model.User,*sessionM.Session, error) {
	otp, err := s.OtpRepo.Get(ctx, otpID)
	if err != nil {
		return nil, nil,err
	}

	// * aqui sabemos que otp no es null

	if time.Now().After(*otp.ExpiresAt) {
		return nil,nil, errors.New("otp-expired")
	}

	if otpCode != *otp.OptCode {
		return nil,nil, errors.New("otp-code-do-not-match")
	}

	// no es necesario por que emitimos dos tokens y cada uno tiene su propi mideware
	// hasta ahora
	// deberia ser verify-sign-up 
	if *otp.Purpose != constants.VerifyEmailOtpPurpose {
		return nil, nil, errors.New("otp-invalid-purpose")
	}

	//usado? - es mejor eliminarlo
	// * De momento solo eliminaremos
	// aqui no usamos el * por que o es el puntero del filed?
	err = s.OtpRepo.Delete(ctx, otp.ID.(string))
	if err != nil {
		return nil, nil, err
	}
	// y actualizar user verificado, debería retornar el user?
	err = s.UserRepo.ConfirmEmail(email)
	if err != nil {
		return nil,nil, fmt.Errorf("confirm-email-error %v", err)
	}
	// * ------------------------------------------ Retornar la session

	// * obtener el usaurio verificado
	user, err := s.UserRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, nil, err
	}

	// como sabes que es password y no google u otro lo mismo para los demás
	// crear para ambos user-not found
	// * verificar si existe el proveedor "email-passwod"
	auth, err := s.AuthRepo.GetByIDProvider(ctx, user.ID.(string), "password")
	if err != nil {
		return nil, nil, err
	}

	// * Assign roles
	roles, err := s.RoleRepo.GetByUserID(ctx, user.ID.(string))
	if err != nil {
		return nil,nil, err
	}

	user.Roles = roles

	// * dos formas instanciar el session entonces y no deberíamos devolver accesstoken y refrestoken
	// * si no asignarle a la session

	// * O enviar los primitivos roles, userID Remenberme al s.SessionSrv.Create
	// ? need entity ?
	// ? agragar comoentidad al

	// []string{"role-test"}

	// mejor pasar los primitivos
	// ! roles test?
	session := &sessionM.Session{
		UserID:    user.ID.(string),
		Email:     user.Email,
		// * variable de entorno?
		RemenberMe: false,
	}

	// ? Deberían estar en Auth struct?
	err = s.SessionSrv.Create(session,roles)
	if err != nil {
		return nil,nil, err
	}

	//user.AuthProviders =append(user.AuthProviders, auth)
	user.AuthProviderCreate = auth
	// * clean the tributtes y en que parte, no me sale en l respuesta pero igul lo voy a limpiar
	//auth.Password = ""
	//auth.PasswordHash = nil

	return user,session, nil
}
