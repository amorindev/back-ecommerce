package service

import (
	"context"
	"errors"
	"time"

	authErr "com.fernando/pkg/app/auth/errors"
	userErr "com.fernando/pkg/app/user/errors"
	userM "com.fernando/pkg/app/user/model"

	"com.fernando/pkg/app/auth/service/utils"
	otpM "com.fernando/pkg/app/otp-codes/model"
)

/*
var (
	ErrUserAlreadyExists = errors.New("user-already-exists")
	)
*/

// es posible que el usuario tenga una cuenta con google por ejemplo por
// eso revisamos el user y el auth con el provider
func (s *Service) SignUp(ctx context.Context, userParam *userM.User, emailOtp bool) (string, error) {
	// * sirve para crear solo el auth y el auth con user
	now := time.Now().UTC()
	provider := "password"
	userParam.AuthProviderCreate.Provider = &provider
	err := userParam.AuthProviderCreate.HashPassword()
	if err != nil {
		return "", err
	}
	// * verificar si existe el usuario
	user, err := s.UserRepo.GetByEmail(ctx, userParam.Email)
	if err != nil {
		if err != userErr.ErrUserNotFound {
			return "", err
		}
	}

	//if user != nil{} user-already-existso auth provider alredyexists

	if user != nil {
		// * verificar si existe el auth "password"
		auth, err := s.AuthRepo.GetByIDProvider(ctx, user.ID.(string), "password")
		if err != nil {
			if err != authErr.ErrAuthNotFound {
				return "", err
			}
		}

		if auth != nil {
			// password-account-alredy-exists
			return "", errors.New("email-already-in-use")
		}

		// ! con vertical slicing  ver cual se va usar stripe no usa otp ver
		// stripe usa tokens?
		if emailOtp {
			code, err := utils.GenOtpCode()
			if err != nil {
				return "", err
			}

			purpose := "verify-email"
			expiresAt := now.Add(time.Hour)
			used := false
			otp := &otpM.OtpCodes{
				//ID: , database
				//AuthID: , en la transaccion
				OptCode:   &code,
				Purpose:   &purpose,
				ExpiresAt: &expiresAt,
				Used:      &used,
				CreatedAt: &now,
			}

			// !asignar el id del user
			err = s.AuthTx.SignUpWithOtp(ctx, userParam.AuthProviderCreate, otp)
			if err != nil {
				return "", err
			}

			err = s.EmailSrv.SendVerificationWithOTP(user.ID.(string), user.Email, code)
			if err != nil {
				return "", err
			}

			// * En este punto deberíamos tener el id
			return otp.ID.(string), nil

		}

		// * Crear solo el auth
		// ! userParam.AuthProviderCreate.UserID = user.ID
		err = s.AuthRepo.Insert(ctx, userParam.AuthProviderCreate)
		if err != nil {
			return "", err
		}

		// * no se si el auth esta traendo el user
		// * Send email verification
		err = s.EmailSrv.SendVerification(user.ID.(string), user.Email)
		if err != nil {
			return "", err
		}

		// cloro aqui es sin otp
		return "", nil
		// !return nil ? para salir de la función a si para login with providers
	}

	// * Crear el usuario
	userParam.EmailVerified = false

	// * assign roles and permisions, verificar si eexisten?
	userRoles := []string{"USER", "USER-TEST"}
	rolesModel, err := s.RoleRepo.GetByNames(ctx, userRoles)
	if err != nil {
		return "", err
	}

	// ? uso append o lo asigno ver comportamiento
	// como menejarlo si es nil
	// o usarauth.UserAgregate.Roles = append(auth.UserAgregate.Roles, role)
	userParam.RolesModel = rolesModel
	// no hay necesida de consutlar GetByNamesLohizo
	userParam.Roles = userRoles

	// ! hacer esta validacion desde aqui? o desde el handler creando un servicio para cada uno
	// o cuando agregar una capa mas?
	if emailOtp {
		code, err := utils.GenOtpCode()
		if err != nil {
			return "", err
		}
		purpose := "verify-email"
		expiresAt := now.Add(time.Hour)
		used := false
		otp := otpM.OtpCodes{
			//ID: , database
			//AuthID: , en la transaccion
			OptCode:   &code,
			Purpose:   &purpose,
			ExpiresAt: &expiresAt,
			Used:      &used,
			CreatedAt: &now,
		}
		err = s.AuthTx.SignUpWithOtpUser(ctx, userParam, &otp)
		if err != nil {
			return "", err
		}

		err = s.EmailSrv.SendVerificationWithOTP(userParam.ID.(string), userParam.Email, code)
		if err != nil {
			return "", err
		}
		// * En este punto deberíamos tener el id
		return otp.ID.(string), nil
	}

	// * Create new account
	err = s.AuthTx.SignUpUser(ctx, userParam)
	if err != nil {
		return "", err
	}

	// ? sacar del
	// * enviar email de confirmacion
	// ? ID del user o auth ? ,
	err = s.EmailSrv.SendVerification(userParam.ID.(string), userParam.Email)
	if err != nil {
		return "", err
	}

	// * limpiar valores que no van a ser mostrados por seguridad, complementar con json omitempty
	//userParam.AuthProviders[0].Password = ""
	//userParam.AuthProviders[0].PasswordHash = nil

	userParam.AuthProviderCreate.Password = ""
	userParam.AuthProviderCreate.PasswordHash = nil

	// authParam.PasswordHash = nil
	return "", nil
}
