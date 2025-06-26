package service

import "com.fernando/pkg/email/service/templates"

type ForgotPasswordTmplData struct {
	Name    string
	Subject string
	Link    string // link de retorno a la app no contendrá token por que usaremos otp
	Code    string
}

func (s *EmailService) SendForgotPassword(userID string, email string, code string) error {
	// * creo el código - me parece buen punto crearlo desde aqui ver

	link, err := s.BranchioAdapter.CreateBranchLink(nil)
	//link, err := s.BranchioAdapter.CreateBranchLink(token)
	if err != nil {
		return err
	}

	data := ForgotPasswordTmplData{
		Name:    "AuthTemplate",
		Subject: email,
		Link:    link,
		Code:    code,
	}

	tmplString, err := templates.LoadTemplate("pkg/email/service/templates/forgot-password-otp.html", data)
	if err != nil {
		return err
	}

	return s.EmailAdapter.Send("test@fernandev.tech", email, tmplString)
}

/* func GenOtpCode2() (string, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return "", err
	}
	n:= nBig.Int64()
	return strconv.Itoa(n.(int)), nil
}
*/
