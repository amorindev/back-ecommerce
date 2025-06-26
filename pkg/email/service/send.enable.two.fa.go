package service

import (
	"com.fernando/pkg/email/service/templates"
)

type EnableTwoFaTmplData struct {
	Name    string
	Subject string
	Code    string
}

func (s *EmailService) SendEnableTwoFa(userID string, email string, code string) error {
	println(email)
	//fmt.Printf("*********-------------*****************-------------")
	data := EnableTwoFaTmplData{
		Name:    "AuthTemplate",
		Subject: email,
		Code:    code,
	}

	tmplString, err := templates.LoadTemplate("pkg/email/service/templates/enable-two-fa.html", data)
	if err != nil {
		return err
	}
	return s.EmailAdapter.Send("test@fernandev.tech", email, tmplString)
}
