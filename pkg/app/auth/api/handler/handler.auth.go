package handler

import (
	"net/http"

	"com.fernando/cmd/api/middlewares"
	authPort "com.fernando/pkg/app/auth/port"
	sessionPort "com.fernando/pkg/app/session/port"
)

// en los handlers verificar el TokenType del token
// y que datos debe tener el token
// que validar segun los tokens de login y os otros junto a los middlewares y expires at

//
// verificar el token en la base de datos
/* if !tokenValido(branchBody.Token) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Invaid token"})
	return
} */
// que se valida en el handler y en el middleware y en refresh token
// las validaciones de token en el middleware handler segun el tipo login on refresh flujo - token email verify u otros
// filtrar por proposito del token y donde ?
// verifaicar el exp desde el frotnend antes de enviar una solicitud, o enviar si tiene un tipo de
// error token invalid hacer una consulta a refresh token handler
// limite de intentos los handlers que lo requieran
type Handler struct {
	AuthService authPort.AuthSrv
	//AuthService  authSrv.AuthService
	SessionService sessionPort.SessionSrv
}

// * me parece que no se usa el *Handler no lo estoy usando desde v1
func NewHandler(server *http.ServeMux, authSrv authPort.AuthSrv, tokenSrv sessionPort.SessionSrv) *Handler {
	h := &Handler{
		AuthService:    authSrv,
		SessionService: tokenSrv,
	}

	// me parece que se puede agregar logguer midleware al mex para no
	// estar agregando a cada uno individual ver en prod que no se use
	// agreupar handlers que necesitan auth

	// TODO cambiar los nobres ver lo que falta erar send email pero debe ser resend
	// TODO el primer envio se realiza desde el servicio  ypor eso rensend

	// ! agregar AuthMiddleware a todos lo que necesitan
	// ! ademas ver el orden de los idlewares

	//! ver que handler son de user de toptcodes u otro

	// ! hice el versionamiento pero no gregue v1
	// que handlers necesitan auth midleware, y autorization middleware

	// ? crear un verify para todos ruta, que dice vertiacal slicing?

	// * Verificar todos los que necesitan midlewares y cuales
	// * Verificar los métodos correctos GET POST

	// * como quitar los midleares de logs en produccion

	signUpH := middlewares.LoggerMiddleware(h.SignUp)
	server.HandleFunc("POST /v1/auth/sign-up", signUpH)

	// ver el limit rating por que es un usuario no atenticado lo smismo para sign in y sign up
	sEVOtpH := middlewares.LoggerMiddleware(h.SendEmailVerificationOtp)
	server.HandleFunc("POST /v1/auth/send-email-verification-otp", sEVOtpH)

	// no importa si lo enviaste por email o sms o sign-up/verify-otp
	signUpVerifyOtpH := middlewares.LoggerMiddleware(h.SignUpVerifyOtp)
	server.HandleFunc("POST /v1/auth/sign-up-verify-otp", signUpVerifyOtpH)

	signInH := middlewares.LoggerMiddleware(h.SignIn)
	server.HandleFunc("POST /v1/auth/sign-in", signInH)

	// * ------------------------------------------------------ start
	// ! como seran  los nombres si es veridicar email ya se save que tiene que ser enviando
	// ! un email si  es verificar telefono se sabe que debe enviarse un telefono
	// ! para otro tipo recuperar contraseña u otro cualquiera ahi si parametriado header
	// ! el tema es con otp o sin otp ver
	// * para enable-2fa-phone se puede hacer dos handlers
	// * uno es para activar desde la pantalla de phones verify phone
	// * y el otro seria al momento de seleccionr como facebook
	// * mediante sms o phone mostrar los telefonos y aahi que el usuario seleccione
	// ? que acciones permiten marcar como verificado ver
	// * voy hacer handler independedientes de momento me enfoco en enable-2fa-phone
	// hay varios tipo con qr con phone ver facebook
	// cmabiar a post cuando tengra varios tipos de como qr
	// MFA parece igual
	// se necesita phoneID
	// ? que pasa en sendSmsVerifyPhoneOtpH,VerifyPhoneOtp,si esta verificado el phone
	// TODO deve ser resend otp sms algo asi // este handler se para enviar el otp 
	// TODO y actualizar el twofa  no para solamente verificar el phone para ello existe otro
	// TODO handler
	resendSmsVerifyPhoneOtpH := middlewares.LoggerMiddleware(middlewares.AuthMiddleware(h.ResendSmsVerifyPhoneOtp))
	server.HandleFunc("POST /v1/auth/resend-verify-phone-otp", resendSmsVerifyPhoneOtpH)

	// no poner en  funcion si es sms o email poner quitar  
	// verify otp enable 2fa phone / ver los nombres para este grupo tambien seria confimr
	enableTwoFaSmsVerifyOtpH := middlewares.LoggerMiddleware(middlewares.AuthMiddleware(h.EnableTwoFaSmsVerifyOtp))
	server.HandleFunc("POST /v1/auth/enable-2fa-sms-verify-otp", enableTwoFaSmsVerifyOtpH)

	// faltaria el paso de verificar al usuario con datos existentes enviando otp al celular
	enableTwoFaSmsH := middlewares.LoggerMiddleware(middlewares.AuthMiddleware(h.EnableTwoFaSms))
	server.HandleFunc("POST /v1/auth/enable-2fa-sms", enableTwoFaSmsH)

	// ver por que sabemos que el sign in no devuelve el tokesn cuando
	// la sesion cuando no verifico su email y cuando esta activo twa
	// asi que aqui no tendre access token ver 
	// es cuando inicia session
	twoFaSmsVerifyOtp := middlewares.LoggerMiddleware(h.TwoFaSmsVerifyOtp)
	server.HandleFunc("POST /v1/auth/two-fa-sms-verify-otp", twoFaSmsVerifyOtp)

	// * ---------------------------------------------------------- end

	refreshTokenH := middlewares.LoggerMiddleware(middlewares.RefreshTokenMiddleware(h.RefreshToken))
	server.HandleFunc("POST /v1/auth/refresh-token", refreshTokenH)

	signOutH := middlewares.LoggerMiddleware(middlewares.RefreshTokenMiddleware(h.SignOut))
	server.HandleFunc("POST /v1/auth/sign-out", signOutH)

	// * Si solo se va enviar el token seria get
	// * de momento sencillo solo refreshtoken enviar y eliminamos cuenta
	// * deberiamos veificar otp, y enviar un token de typo
	// * hay dos tipos dever facebook , descativar cuenta y eliminar cuenta
	// * facebook usa password - ver cuando usar otp cuando solo enlace como stripe
	// * ver flujos
	deleteAccountH := middlewares.LoggerMiddleware(middlewares.AuthMiddleware(h.DeleteAccount))
	server.HandleFunc("POST /v1/auth/delete-account", deleteAccountH) // ?POST

	// agregar sms al nombre - este handler solo seria para verificar
	// esto no se si va en user o en auth o en phones ver DDD
	verifyPhoneOtpH := middlewares.LoggerMiddleware(middlewares.AuthMiddleware(h.VerifyPhoneOtp))
	server.HandleFunc("POST /v1/auth/verify-phone-otp", verifyPhoneOtpH)
	// * --------------------------------------

	//server.HandleFunc("POST /v1/auth/send-sms-verification", )

	server.HandleFunc("POST /v1/auth/send-email-verification", h.SendEmailVerification)
	// puede ser sign-up/verify y poner en el body si es emil o sms, lo que pasa
	server.HandleFunc("GET /v1/auth/sign-up-verify-email", middlewares.AuthMiddleware(h.VerifyEmail)) // es get?

	server.HandleFunc("POST /v1/auth/forgot-password", h.ForgotPassword)
	// deberia ser en header el token?, si es asi solo seria get
	server.HandleFunc("POST /v1/auth/verify-forgot-password", middlewares.AuthMiddleware(h.VerifyForgotPassword))
	// si
	server.HandleFunc("POST /v1/auth/change-password", h.ChangePassword)                      // si
	server.HandleFunc("PUT /v1/auth/new-password", middlewares.AuthMiddleware(h.NewPassword)) //si

	// envias el otp mediante sms
	//server.HandleFunc("POST /v1/auth/sign-up-phone", )

	// Login with providers ,
	// ? oauth ?
	server.HandleFunc("POST /v1/auth/google", h.GoogleSignIn)
	server.HandleFunc("POST /v1/auth/apple", h.AppleSignIn)
	server.HandleFunc("POST /v1/auth/facebook", h.FacebookSignIn)

	// * Current user es diferente a auth changes

	// ifgual a firebase ses o stream
	// auth changes igual a firebase  ssr, o websockets
	// roles y permisiones?
	server.HandleFunc("GET /v1/auth/auth-changes", h.AuthChanges) // si

	// cerrar sesiones despues de cambiar el password es un handler aparte
	// o se integrta en el servicio de new-password
	// si va pasar de register a send email verification y no al login deberia generar el token ahi
	// en verify.email.go
	return h
}
