package service

import (
	"context"
	"time"

	"com.fernando/pkg/app/auth/service/constants"
	"com.fernando/pkg/app/auth/service/utils"
	"com.fernando/pkg/app/otp-codes/model"
)

// * limit rating tanto en el frontend como en el backend
// * Por ejemplo en el handler verificar que  solo afecte a ese usaurio necesitas .env?
// * para los rating limints
// * en el caso de mobile Visible cambialo a false puedes intentar desde de 30seg
// * como prueva, por que en el frontend para no hacer peticiones innecesarias
func (s *Service) SendEmailVerificationOtp(ctx context.Context, email string) error {
	// * segun el flujo no es necesari verificar si existe el user
	// * por que se supone si el sign up es exito esta creado
	// * y si va desde la pantalla de sign in tambien verifica si existe
	// * si por que puede ser otroofrontend  con otro flujo de momento ir validando la mayor parte si existe el usuario
	code, err := utils.GenOtpCode()
	if err != nil {
		return err
	}

	// * devería ser verify-account o algo asi
	now := time.Now()
	purpose := constants.VerifyEmailOtpPurpose
	expiresAt := now.Add(time.Hour)
	used := false

	user, err := s.UserRepo.GetByEmail(context.Background(), email)
	if err != nil {
		return err
	}

	otp := &model.OtpCodes{
		//ID: , base de datos
		UserID:    user.ID.(string),
		OptCode:   &code,
		Purpose:   &purpose,
		ExpiresAt: &expiresAt,
		Used:      &used,
		CreatedAt: &now,
	}

	err = s.OtpRepo.Insert(context.Background(), otp)
	if err != nil {
		return err
	}

	err = s.EmailSrv.SendVerificationWithOTP(user.ID.(string), email, code)
	if err != nil {
		return err
	}

	return nil
}
