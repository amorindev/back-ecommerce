package resend

import (
	"com.fernando/pkg/email/port"
	"github.com/resend/resend-go/v2"
)

var _ port.EmailAdapter = &ResendAdapter{}

type ResendAdapter struct {
	Client *resend.Client
}

func NewAdapter(client *resend.Client) *ResendAdapter{
	return &ResendAdapter{
		Client: client,
	}
}