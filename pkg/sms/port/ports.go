package port


type SmsAdp interface{
    // menadaje o plantilla coom el email?
    // el codigo es parte de la logica de servicio 
    // adapter se enfoca en lo que ofrecen los servicios
    Send(from, to, msg string) error
}

// ! falta ver el context en los serciios
// * ver si se puede agregar para verificar telefono, correo con otp o sin otp
type SmsSrv interface {
    // ver el nombre o nombres de los servicios
    // le pasamos el code de otp ver si genrar el otp y guardar en la base de datos
    // iria dentro de SmsSrv y no donde esta 
  SendVerificationOtp(to string, code string) error
}


type EmailSrv interface {
	SendEmailVerification(userID string, email string) error
	SendEmailVerificationWithOTP(userID string, email string, code string) error
	SendForgotPassword(userID string, email string, code string) error
}
