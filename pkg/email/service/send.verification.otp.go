package service

import (
	"com.fernando/pkg/email/service/templates"
)

type ForgotPasswordOTPTmplData struct {
	Name    string
	Subject string
	Link    string // link de retorno a la app no contendrá token por que usaremos otp
	Code    string
}

func (s *EmailService) SendVerificationWithOTP(userID string, email string, code string) error {

	link, err := s.BranchioAdapter.CreateBranchLink(nil)
	if err != nil {
		return err
	}

	data := ForgotPasswordOTPTmplData{
		Name:    "AuthTemplate",
		Subject: email,
		Link:    link,
		Code:    code,
	}

	//pkg/email/service/templates/email-verification-otp.html
	tmplString, err := templates.LoadTemplate("pkg/email/service/templates/email-verification-otp.html", data)
	if err != nil {
		return err
	}

	return s.EmailAdapter.Send("test@fernandev.tech", email, tmplString)
}
