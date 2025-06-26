package port


// * en los adaters se enfocan en los servicios que tienen los seridores de email como send email u otro
//to string, subject string, body string
type EmailAdapter interface {
	Send(from, to, htmlBody string) error
	//Verification(from, to, htmlBody string) error
	ForgotPassword() error // esto no seria por que parece logica de negocio va en el ser
}

// SendEmailVerification ya se sabe que es email
// * Me parece que el userID era para branchio
type EmailSrv interface {
	// * ver si con tokens es inseguro o el problema de validar el purpose se puede cometer errores si olvidamos
	// * validarlo en la request si no eliminarlo y solo trabajamos con tokens para la sesion
	SendVerification(userID string, email string) error
	// el codigo me parece que se genera en el sercicio aqui no en se pasa como parametro
	SendVerificationWithOTP(userID string, email string, code string) error
	SendForgotPassword(userID string, email string, code string) error
	SendEnableTwoFa(userID string, email string, code string) error
	SendTwoFaSignIn(userID string, email string, code string) error
}
