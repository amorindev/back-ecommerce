package core

import (
	authM "com.fernando/pkg/app/auth/model"
	sessionM "com.fernando/pkg/app/session/model"
	userM "com.fernando/pkg/app/user/model"
)

// * definir que request structuras dentro del handler (no es ideal), que struct
// * se pueden estar reutilizando para cada request

// * cada estructura se agregará a su vertical slicing

// ?se deveria agregar una entidad mas, tenemos 2
// ? se valida si User == nil ? falta datos del usaurio - revisar flujo

/* type User struct{} */

// hacer login debe retornar is gerified  or currentuser?
// para mostrar la página de send email
// * si creamos el sign in response ya no necesitamos agregar acess ni refresh al Auth
// * se podría sign in response usar un puntero *User y solo pasarle el puntero
// * el vervicio retornaría el usar el access y refresh por aparte ver cual es mejor

// * Update password mediante email
type NewPasswordRequest struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// * Update password sabiendo su actual contraseña
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password"` // required min 2
	Password        string `json:"password"`         //required min 2
	ConfirmPassword string `json:"confirm_password"` // required min 2
}

type ForgotPasswordReq struct {
	Email string `json:"email"` // en este caso si es necesario en el body por que no tien el token
}

type VerifyEmailOTP struct {
	// ver por que la ligca es  no dar  lasessio n hasta que verifique su cuenta
	// ! claro aqi el problema era que al realizar sign up no retornavamos la session si no
	// ! al verificar el email ahora es tiempo de refactorizar y despes delsign up retornar
	// ! la session y al verificar VerifyEmailOTP tendremos el userID o email
	// ! para marcar como verificado
	OtpID   string `json:"otp_id"`
	OtpCode string `json:"otp_code"`
	Email   string `json:"email"`   // se deve sacar del token
	UserID  string `json:"user_id"` // ! todaía no se usa
	// ahi es email o userID para actualizar a emailverified y no ambos
	// ademas los datos se saca de token
}

// ? que pasa en vertical slicing si tanto sin otp y con otp usan este mismo poner en una
// ? carpeta por encima como los clientes mongo y otros, aunque son diferentes, o
// ? crear uno para cada uno, los clients como conexion a base de datos tienen mas peso
// ? que pasa con resend client o branhcio que solo estan en un sercio
// ? como deerían conectarse cada servicio y dependecias
// Cuando el usario se aha registrado no se valída si exiteste claro,
// pero cuando
// estas structuras pequeñas se podrían crenr en la misma función?
type SendEmailVerificationReq struct {
	Email string `json:"email"`
}

// como estructurar el auth estructura con use estructura
// junto para crear el registro separado
// con un trrigger para la base de datos o  un transaccion

type AuthEntity struct {
	Email    string
	Password string
	//User     User
}

// * Recuera User con auth son uno a muchos entonces para creareel auth solo necesito el userID como microblog

type RefreshTokenResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	//ExpiresIn    int64  `json:"expires_in"`
}

type AuthResp struct {
	User        *userM.User       `json:"user"`
	Session     *sessionM.Session `json:"session"`
	Credentials *authM.Auth       `json:"credentials"`
	OtpID       string            `json:"otp_id,omitempty" bson:"-"` // * aux Response
}
