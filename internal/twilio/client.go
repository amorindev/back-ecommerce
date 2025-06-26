package twilio

import (
	"github.com/twilio/twilio-go"
)

func NewClient() *twilio.RestClient{
    // twilio obtendra las api key por nosotros si seguimos su
    // con el nombre que maneja dentro del paquete TWILIO...
    return twilio.NewRestClient()
}
