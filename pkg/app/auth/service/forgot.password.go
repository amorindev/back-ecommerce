package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"com.fernando/internal/claim"
	"com.fernando/pkg/app/auth/service/utils"
	"com.fernando/pkg/app/otp-codes/model"
)

func (s *Service) ForgotPassword(ctx context.Context, email string) error {
	// requerimientos
	// que exista el email en la base de datos

	u, err := s.UserRepo.GetByEmail(context.Background(), email)
	if err != nil {
		// seria user-not-fund
		return err
	}

	// * Si el auth no estaría dentro de user no tendríamos este proceso
	a, err := s.AuthRepo.GetByIDProvider(ctx, u.ID.(string), "password")
	if err != nil {
		// seria user-not-found
		return err
	}

	if a == nil {
		return errors.New("user-not-found")
	}

	// create token
	idStr, ok := u.ID.(string)
	if !ok {
		return errors.New("auth service - user ID no pudo ser convertido a string")
	}

	c := claim.NewForgotPasswordToken(idStr, u.Email)

	// os.getenv or config ?
	config, err := claim.GetConfig()
	if err != nil {
		return err
	}

	// tokenString
	_, err = c.GetToken(config.AccessString)
	if err != nil {
		return fmt.Errorf("forgotPassword serv err - getToken: %w", err)
	}

	code, err := utils.GenOtpCode()
	if err != nil {
		return err
	}

	purpose := "forgot-password"
	used := false
	now := time.Now()
	expiresAt := now.Add(time.Hour * 2)
	// guardar en la base de datos
	otp := &model.OtpCodes{
		UserID:    u.ID,
		OptCode:   &code,
		Purpose:   &purpose,
		Used:      &used,
		ExpiresAt: &expiresAt,
		CreatedAt: &now,
	}
	err = s.OtpRepo.Insert(ctx, otp)
	if err != nil {
		return err
	}

	// enviar el corrreo con un token de expiracion sera validado en new-password handler
	err = s.EmailSrv.SendForgotPassword(u.ID.(string), u.Email, code)
	if err != nil {
		return err
	}
	return nil
}
