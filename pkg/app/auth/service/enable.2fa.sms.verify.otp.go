package service

import (
	"context"
	"errors"
	"time"

	authCts "com.fernando/pkg/app/auth/service/constants"
	userCts "com.fernando/pkg/app/user/constants"
	"com.fernando/pkg/app/user/model"
)

func (s *Service) EnableTwoFaSmsVerifyOtp(ctx context.Context, otpID string, otpCode string, userID string) (*model.User, error) {
	
	otp, err := s.OtpRepo.Get(ctx, otpID)
	if err != nil {
		return nil, err
	}

	if time.Now().After(*otp.ExpiresAt) {
		return nil, errors.New("otp-expired")
	}

	if otpCode != *otp.OptCode {
		return nil, errors.New("otp-code-do-not-match")
	}

	if *otp.Purpose != authCts.VerifyPhoneOtpPurpose {
		return nil, errors.New("otp-invalid-purpose")
	}
	//usado? - es mejor eliminarlo
	// * De momento solo eliminaremos
	// aqui no usamos el * por que o es el puntero del filed?

	err = s.OtpRepo.Delete(ctx, otp.ID.(string))
	if err != nil {
		return nil, err
	}
	twoFaMethod := userCts.MethodSms
	// user two fa enable y userTwoFaSmsID confirmed
	// que pasa si cancela en algunos pasos como en el primer handler
	// entonces quedaria en verifcar o como sería
	// seria una transaccion de momento simple
	// actualizar el campo TwoFaMethod
	err = s.UserRepo.EnableTwoFaSms(ctx, userID, twoFaMethod)
	if err != nil {
		return nil, err
	}
	// falta actualizar  userTwoFa la otra colleccion

	// * Retornar User de momento, por que afecta solo al user y como creamos una nueva session
	// * en sign in si sabemos que es password pero aqui no sabríamos que porveedor
	// * es anocer que lo pasemos, otra forma es recuperando la session
	// * siguiendo la recomendacion retornar solo si de chat
	user, err := s.UserRepo.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	// * Assign roles
	roles, err := s.RoleRepo.GetByUserID(ctx, user.ID.(string))
	if err != nil {
		return nil, err
	}

	user.Roles = roles

	return user, nil
}
