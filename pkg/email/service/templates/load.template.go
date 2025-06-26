package templates

import (
	"bytes"
	"fmt"
	"html/template"
)

func LoadTemplate(templatePath string, data any) (string, error) {
	// Parse files acepta las imagenes?
	// agraga imagen al html
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("parse files - loadTemplate: %w", err)
	}

	var body bytes.Buffer

	// crearlos headers desde aqui?

	/* err = t.Execute(&body, ForgotPasswordTmplData{
		Name:    appName,
		Subject: subject,
		Link:    link,
		Code:    verificationCode,
	}) */

	err = t.Execute(&body, data)

	if err != nil {
		return "", fmt.Errorf("execute - loadtemplate err: %w", err)
	}

	return body.String(), nil
}
