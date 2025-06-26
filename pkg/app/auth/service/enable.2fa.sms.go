package service

import (
	"context"
	"errors"
	"time"

	"com.fernando/pkg/app/auth/service/constants"
	"com.fernando/pkg/app/auth/service/utils"
	otpModel "com.fernando/pkg/app/otp-codes/model"
)

// ! que verificaciones se necesita en todos los handler
// ! en especial enble2fa sms y el verify que pasa si ya esta activo
// ! se puede actualizar er facebook mas adelante
// ! ver que se va empviar al frontend por ejemplo se puede agregar el methodo
// ! para desirle al usuario que metodo agregado dese cambiar el teléfono o algo asi
// * ver como es cuando no existe el phone como facebook crear otr página
// * ver el flujo de paginas tambien
func (s *Service) EnableTwoFaSms(ctx context.Context, userID string, phoneID string) (string, error) {
	// * verificar si existe el phone
	/* phone, err := s.PhoneRepo.Get(context.Background(), phoneID)
	// si no existe tambien lo va retornar error
	if err != nil {
		return "", err
	} */

	// * insertar twofaauth collection
	/* twoFaSms := model.NewUserTwoFaSms(userID, phoneID, false)
	err = s.UserRepo.InsertTwoFaSms(ctx, twoFaSms)
	if err != nil {
		return "", err
	} */
	user, err := s.UserRepo.Get(ctx, userID)
	if err != nil {
		return "", err
	}

	if user.IsTwoFaEnabled {
		return "", errors.New("authentication two fa habilitado")
	}
	// * retornar el codeID

	code, err := utils.GenOtpCode()
	if err != nil {
		return "", err
	}
	purpose := constants.VerifyPhoneOtpPurpose
	now := time.Now()
	expiresAt := now.Add(time.Hour)
	used := false
	otp := &otpModel.OtpCodes{
		//ID: , database
		//AuthID: , en la transaccion
		OptCode:   &code,
		Purpose:   &purpose,
		ExpiresAt: &expiresAt,
		Used:      &used,
		CreatedAt: &now,
	}

	otp.UserID = userID

	err = s.OtpRepo.Insert(ctx, otp)
	if err != nil {
		return "", err
	}

	// * De momento usaremos email twilio cobra bien
	// ! usar el phone que se relaciona con userPhonetwofa
	// ! sin importar si esta verificado
	/* fmt.Printf("Phone number service %s\n", *phone.Number)
	err = s.SmsSrv.SendVerificationOtp(*phone.Number,*otp.OptCode)
	if err != nil {
		return "", err
	} */
	err = s.EmailSrv.SendEnableTwoFa(userID, user.Email, code)
	if err != nil {
		return "", err
	}

	// ? marcar phone como verificado el phone si no lo esta como hacerlo
	// ? crear metdata i guardar ewl phone id en el otp? ver flujos
	return otp.ID.(string), nil
}
