package adapter

import (
	"com.fernando/pkg/app/auth/port"
	"github.com/coreos/go-oidc/v3/oidc"
)

var _ port.AuthAdapter = &Adapter{}

// GoogleProvider ?
type Adapter struct {
	GoogleVerifier *oidc.IDTokenVerifier
	AppleProvider  *AppleProvider
}

// appleProvider *AppleProvider
/* func NewAdapter(googleVerifier *oidc.IDTokenVerifier) *Adapter {
	return &Adapter{
		GoogleVerifier: googleVerifier,
		//AppleProvider:  appleProvider,
	}
}
 */

func NewAdapter() *Adapter {
	return &Adapter{
		
		//AppleProvider:  appleProvider,
	}
}