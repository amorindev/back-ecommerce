package resend

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

// go embed no admite tmpls fuera del paquete si deseas utilizaro en otro
// lugar deberás copiar el archivo

// dgo:embed tmpls/*
var templateFS embed.FS

type EmailData struct {
	AppName string
	Subject string
}

var emailData EmailData

// ! me parece que la loogica de aqui es compatible con otro proveedor de email
func (a *ResendAdapter) ForgotPassword() error {
	// recipient string, templateFile string, data EmailData

	// templateFile
	absolutePath := filepath.Join("tmpls", "forgotpassword.html")

	//template.ParseFiles("")

	tmpl, err := template.ParseFS(templateFS, absolutePath)
	if err != nil {
		return fmt.Errorf("ForgotPassword - ResendAdapter: %w", err)
	}

	emailData.AppName = os.Getenv("APP_NAME")

	subject := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(subject, "subject", emailData)
	if err != nil {
		return fmt.Errorf("forgot password - ResendAdapter: %w ", err)
	}

	return nil
}

func (a *ResendAdapter) ForgotPassword2() error {

	return nil
}
