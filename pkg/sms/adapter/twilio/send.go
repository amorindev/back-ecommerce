package twilio

import (
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func (a *Adapter) Send(from, to, msg string) error {
	params := &api.CreateMessageParams{}
	params.SetBody(msg)
	params.SetFrom(from)
	params.SetTo(to)
	// la respuesta lo vamos a omitir de momento
	_, err := a.TwilioClient.Api.CreateMessage(params)
	if err != nil {
		return err
	}
	return nil
}
