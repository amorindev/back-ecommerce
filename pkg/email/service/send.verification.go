package service

import (
	"com.fernando/internal/claim"
	"com.fernando/pkg/email/service/templates"
)

// para los tokens time
//agregar imagenes - variables de entorono
//domain
// si se incluyen inagenes no agregar al .dockerignore, por que serán llamadas desde go

// definir las dos formas de generar el body html ventajas y escojer uno

type EmailVerificationData struct {
	Name    string
	Subject string
	Link    string
}

// * usar el token service
const key = "token"

// ! me parece que no debería estar la creacion de token aqui si no service auth sign up
func (s *EmailService) SendVerification(userID string, email string) error {
	// debe ser jwt token? deberái haber otro hander para validarlo
	c := claim.NewEmailVerificationToken(userID, email)

	config, err := claim.GetConfig()
	if err != nil {
		return err
	}

	token, err := c.GetToken(config.AccessString)
	if err != nil {
		return err
	}

	linkData := map[string]string{
		key: token,
	}

	link, err := s.BranchioAdapter.CreateBranchLink(linkData)
	if err != nil {
		return err
	}

	/* tmplString, err := loadTemplateEmailVerification("Yanbal", "fernando", link)
	if err != nil {
		return err
	} */

	tmplData := EmailVerificationData{
		Name:    "AuthTemplate",
		Subject: email,
		Link:    link,
	}

	tmplString, err := templates.LoadTemplate("pkg/email/service/templates/email-verification.html", tmplData)
	if err != nil {
		return err
	}

	// ? domain en variables de entorno?
	//err = s.EmailAdapter.Verification("Acme <onboarding@resend.dev>", "fernandoamorindev@gmail.com", tmplString)
	err = s.EmailAdapter.Send("test@fernandev.tech", email, tmplString)
	if err != nil {
		return err
	}

	//go s.EmailAdapter.Verification(link)

	//save the token
	return nil
}
