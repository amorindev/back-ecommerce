package service

import "com.fernando/pkg/email/service/templates"

type TwoFaSignInTmplData struct {
	Name    string
	Subject string
	//Link    string // link de retorno a la app no contendrá token por que usaremos otp
	Code string
}

func (s *EmailService) SendTwoFaSignIn(userID string, email string, code string) error {
	data := ForgotPasswordTmplData{
		Name:    "AuthTemplate",
		Subject: email,
		Code:    code,
	}
	tmplString, err := templates.LoadTemplate("pkg/email/service/templates/two-fa-sign-in.html", data)
	if err != nil {
		return err
	}

	return s.EmailAdapter.Send("test@fernandev.tech", email, tmplString)
}
