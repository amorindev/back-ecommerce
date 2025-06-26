package resend

import (
	"github.com/resend/resend-go/v2"
)

// to, subject, body string
func (a *ResendAdapter) Send(from, to, htmlBody string) error {
	params := &resend.SendEmailRequest{
		From:    from,
		To:      []string{to},
		Subject: "Hello",
		Html:    htmlBody,
		/* Headers: map[string]string{
			"MIME-Version": "1.0",
			"Content-Type": "text/html; charset=UTF-8",
		}, */
	}

	_, err := a.Client.Emails.Send(params)
	if err != nil {
		return err
	}
	return nil
}
