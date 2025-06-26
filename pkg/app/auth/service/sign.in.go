package service

import (
	"context"
	"fmt"

	"time"

	"com.fernando/pkg/app/auth/service/constants"
	"com.fernando/pkg/app/auth/service/utils"
	otpModel "com.fernando/pkg/app/otp-codes/model"

	sessionM "com.fernando/pkg/app/session/model"
	userM "com.fernando/pkg/app/user/model"
)

// ? en login se retorna el user?, con la sesion - utilizar las tags json
// ? se debería sacar token y refrehtoken de la auth structure no afect tanto
// ? cuidr los tags
// user-not-found
// auth-not-found

/*
var (
	ErrInvalidCredentials = errors.New("invalid-credentials") // save to auth service
)
*/
// get by emai solo auth o tambien user y si es embedding?
// como funcion el current user con id email? y que datos
func (s *Service) SignIn(ctx context.Context, email string, password string, rememberMe bool) (*userM.User, *sessionM.Session, string, error) {
	// TODO: verificar que rerfesh token se guarde en la base de datos
	// * Verificar si el usuario existe
	user, err := s.UserRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, nil, "", err
	}

	// crear para ambos user-not found
	// * verificar si existe el proveedor "email-passwod"
	auth, err := s.AuthRepo.GetByIDProvider(ctx, user.ID.(string), "password")
	if err != nil {
		return nil, nil, "", err
	}

	// invalid credential donde?
	_, err = auth.PasswordMatch(password)
	if err != nil {
		return nil, nil, "", fmt.Errorf("password not set: %w", err)
	}

	// * Assign roles
	roles, err := s.RoleRepo.GetByUserID(ctx, user.ID.(string))
	if err != nil {
		return nil, nil, "", err
	}

	user.Roles = roles

	// * dos formas instanciar el session entonces y no deberíamos devolver accesstoken y refrestoken
	// * si no asignarle a la session

	// * O enviar los primitivos roles, userID Remenberme al s.SessionSrv.Create
	// ? need entity ?
	// ? agragar comoentidad al

	// mejor pasar los primitivos
	// ! roles test?
	// ? debemos crear en token aqui o en session service ? pasandole solo primitivos, cuando desacoplar
	// y nos retorne session puntero
	session := &sessionM.Session{
		Email:      user.Email,
		UserID:     user.ID.(string),
		RemenberMe: rememberMe,
	}

	// ? Deberían estar en Auth struct?
	err = s.SessionSrv.Create(session, roles)
	if err != nil {
		return nil, nil, "", err
	}

	// * clean the tributtes y en que parte, no me sale en l respuesta pero igul lo voy a limpiar
	auth.Password = "" // como los tags son - pero igual
	auth.PasswordHash = nil
	// user.AuthProviders = append(user.AuthProviders, auth)
	user.AuthProviderCreate = auth

	// * Lo necesitamos para las dos acciones de abajo ver el flujo me para que solo uno se cumple
	code, err := utils.GenOtpCode()
	if err != nil {
		return nil, nil, "", err
	}
	// hasta ahora estamos usando solo primitivos de la Session{}

	// ! verificar si es otp? mejor definir cual se va a usar de momento otp sin validar
	if !user.EmailVerified {
		// ! definir las purposes que tendrá creo que para sign-in y sign-up
		// seria solo veriry-email y por que para validarlo seria si es signin osignup
		// mejor solo verify
		purpose := constants.VerifyEmailOtpPurpose
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

		// * me parece mejor auth tabla pero buenmejor de user
		// * la session si me parece con auth relacion del id ver
		otp.UserID = user.ID

		err = s.OtpRepo.Insert(ctx, otp)
		if err != nil {
			return nil, nil, "", err
		}

		err = s.EmailSrv.SendVerificationWithOTP(user.ID.(string), user.Email, code)
		if err != nil {
			return nil, nil, "", err
		}

		return user, session, otp.ID.(string), nil
	}
	// segun nuestro flujo es necesario que el usuario verifique su email para iniciar session
	// y luego activar el verificacion por email
	// * ver que mas se va a verificar por que en UserTwoFaSms tambien hay una propiedad confirmed
	if user.IsTwoFaEnabled {
		purpose := constants.VerifyPhoneSignInOtpPurpose
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

		// * me parece mejor auth tabla pero buenmejor de user
		// * la session si me parece con auth relacion del id ver
		otp.UserID = user.ID

		err = s.OtpRepo.Insert(ctx, otp)
		if err != nil {
			return nil, nil, "", err
		}

		// ! aqui esta mal por que debe ser el phone de UserTwoFaSms
		// desde donde verificar si el phone es valido
		// desde aqui o desde  phonesrv ver si ver ver de momento aqui
		// tambien ver el flujo por que si llega a este punto deberia tener
		// un telefono ya verificado pero igual demomento tambien lo haré  aqui
		/* if user.Phone == nil {
			return nil,nil,"", errors.New("phone no hay")
		}
		if user.PhoneVerified {
			return nil, nil, "", errors.New("phone is not verified")
		} */
		/* phone, err := s.UserRepo.GetPhonetwoFaSms(context.Background(), user.ID.(string))
		if err != nil {
			return nil, nil, "", err
		}

		// ! usar el phone que se relaciona con userPhonetwofa
		err = s.SmsSrv.SendVerificationOtp(*phone.Number, *otp.OptCode)
		if err != nil {
			return nil, nil, "", err
		} */
		err := s.EmailSrv.SendTwoFaSignIn(user.ID.(string), email, code)
		if err != nil {
			return nil, nil, "", nil
		}
		return user, session, otp.ID.(string), nil
	}
	// * Me parece que es mejor separarlo en otro handler
	// * no deberia tener  aqui una session deberia tenerlo 
	// * si ha verificado su email  que pasa en el twofa
	return user, session, "", nil
}
