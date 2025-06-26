package templates

/* import (
	"bytes"
	"fmt"
	"html/template"
)
 */
/* type TmplData struct {
	Name    string
	Subject string
	Link    string
}

func loadTemplateEmailVerification(appName, subject, link string) (string, error) {
	// Parse files acepta las imagenes?
	// agraga imagen al html
	t, err := template.ParseFiles("pkg/services/email/service/send-email-verification.html")
	if err != nil {
		return "", fmt.Errorf("parse files - loadTemplate: %w", err)
	}

	var body bytes.Buffer

	// crearlos headers desde aqui?

	err = t.Execute(&body, TmplData{
		Name:    appName,
		Subject: subject,
		Link:    link,
	})

	if err != nil {
		return "", fmt.Errorf("execute - loadtemplate err: %w", err)
	}

	return body.String(), nil
}

*/
 