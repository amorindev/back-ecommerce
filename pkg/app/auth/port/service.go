package port

import (
	"context"

	authM "com.fernando/pkg/app/auth/model"
	sessionM "com.fernando/pkg/app/session/model"
	userM "com.fernando/pkg/app/user/model"
)

// * evaluar si sacar access y refreshtoken
// se podría sacar el  token y refrestoken del uauth y solo retornarlo, pero
// entonces no tedría tgs json per si en el vervicio eliminar campo password
// ! verificar todo lo que se va a pasar a user
// ! cuando sacar por separa y cuando junto por ejemplo la sesison pertenece a un suaurio
// ! pero esta separado en la interfaz ver como se vaorganizar, igualmente con el authProvider
// ! y otros subdominios
type AuthSrv interface {
	// create es el sign up Create(ctx context.Context, user *model.User) error
	SignUp(ctx context.Context, userParam *userM.User, emailOtp bool) (string, error) // o withCode
	// otp id es el string
	SignIn(ctx context.Context, email string, password string, remenberMe bool) (*userM.User, *sessionM.Session, string, error)
	// verificar que funciones correctemene get auth por que depende del user
	// retornar el otpID o crear una entidad mas grande qeu abarque todo?
	SendEmailVerificationOtp(ctx context.Context, email string) error
	// esta mal el sign up retorna la session y el user
	SignUpVerifyOTP(ctx context.Context, otpID string, otpCode string, userID string) (*userM.User, *sessionM.Session, error)
	// * o llamarlo signout
	RevokeToken(ctx context.Context, tokenID string) error
	DeleteAccount(ctx context.Context, userID string, password string) error
	// hsta sign out la primera vesion del template
	ForgotPassword(ctx context.Context, email string) error
	NewPassword() error
	ChangePassword() error
	ConfirmEmail(userID string) error // no seria email string
	// cuando cambia contraseña doc del handler con Fern y go senior
	CloseSessions() error
	//Reboked() error
	// google y apple deben tener sercbice o defrente adapter. iniciar seesion y registrar
	GoogleSignIn(ctx context.Context, rememberMe bool, providerTokenID string) (*authM.Auth, error)
	AppleSignIn() error
	FacebookSignIn() error

	// aparrte de esto que otrs puede haver?
	// confiremail en auth user por que devolvera el user no e auth, o tiene que develover auth?
	// y desde el frontend sesion.User = user

	// * Enable 2fa phone debe verificar si hay el phone existe y esta verificado
	// userID para actualizar el  campo enable sign in
	// verificar si se debe agregar algo mas por ejemplo al user agregarle se activo
	// 2fa mediante email, y como seria con MFA una lista
	// se podria marcar como verificado si realiza esta accion el usuario? me parece que si
	// ver los momentos en enviar otp para realizar la acccion o si el
	// usuario quiere usar un phone que no ha sido verificado
	// para ello filtrar por verificado
	// * 1. antes de activar auth 2 factor enviar otp
	// * 2. ver si enviar el otp viendo si se ha verificado o no su phone o da igual enviar nomas
	// * 	se deberia crear un handler para cada tipo de otp me parece que si
	// *	y en el otp marcar como verifcado o solo validar?
	EnableTwoFaSms(ctx context.Context, userID string, phoneID string) (string, error)
	// userTwoFaSmsID para actualizar  a confirmado
	// ver de momento no lo vamos a usar carefull userTwoFaSmsID string
	EnableTwoFaSmsVerifyOtp(ctx context.Context, otpID string, otpCode string, userID string) (*userM.User, error)

	TwoFaSmsVerifyOtp(ctx context.Context, otpID string, otpCode string) (*userM.User, *sessionM.Session, error)
}
