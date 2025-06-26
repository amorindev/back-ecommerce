package twilio

import (
	"com.fernando/pkg/sms/port"
	"github.com/twilio/twilio-go"
)

var _ port.SmsAdp = &Adapter{}

type Adapter struct {
	TwilioClient *twilio.RestClient
}

func NewTwilioAdapter(client *twilio.RestClient) *Adapter {
	return &Adapter{
		TwilioClient: client,
	}
}
