package service

import (
	"context"
	"errors"
	"time"

	"com.fernando/pkg/app/auth/service/constants"
	sessionM "com.fernando/pkg/app/session/model"
	userM "com.fernando/pkg/app/user/model"
)

func (s *Service) TwoFaSmsVerifyOtp(ctx context.Context, otpID string, otpCode string) (*userM.User, *sessionM.Session, error) {
	// esta pendiennte verificar si no existe manejo de erro
	otp, err := s.OtpRepo.Get(ctx, otpID)
	if err != nil {
		return nil, nil, err
	}

	// * aqui sabemos que otp no es null

	if time.Now().After(*otp.ExpiresAt) {
		return nil, nil, errors.New("otp-expired")
	}

	// verificar la propiedad used

	if otpCode != *otp.OptCode {
		return nil, nil, errors.New("otp-code-do-not-match")
	}

	if *otp.Purpose != constants.VerifyPhoneSignInOtpPurpose {
		return nil, nil, errors.New("otp-invalid-purpose")
	}

	//usado? - es mejor eliminarlo
	// * De momento solo eliminaremos
	// aqui no usamos el * por que o es el puntero del filed?
	// ! mas abajo lo estamos usando
	/* err = s.OtpRepo.Delete(ctx, otp.ID.(string))
	if err != nil {
		return nil, nil, err
	} */

	// * obtener el usaurio verificado
	user, err := s.UserRepo.Get(ctx, otp.UserID.(string))
	if err != nil {
		return nil, nil, err
	}

	// como sabes que es password y no google u otro lo mismo para los demás
	// desde la request me parece que debe venir
	// crear para ambos user-not found user y auth not  fund
	// * verificar si existe el proveedor "email-passwod"
	auth, err := s.AuthRepo.GetByIDProvider(ctx, user.ID.(string), "password")
	if err != nil {
		return nil, nil, err
	}

	// * Assign roles
	roles, err := s.RoleRepo.GetByUserID(ctx, user.ID.(string))
	if err != nil {
		return nil, nil, err
	}

	user.Roles = roles

	session := &sessionM.Session{
		UserID: user.ID.(string),
		Email:  user.Email,
		// su constructor de session por defecto false
		RemenberMe: false,
	}

	// ? Deberían estar en Auth struct?
	err = s.SessionSrv.Create(session, roles)
	if err != nil {
		return nil, nil, err
	}

	user.AuthProviderCreate = auth
	// * clean the tributtes y en que parte, no me sale en l respuesta pero igul lo voy a limpiar
	//auth.Password = ""
	//auth.PasswordHash = nil
	return user, session, nil
}
