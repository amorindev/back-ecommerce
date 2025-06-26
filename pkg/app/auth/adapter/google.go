package adapter

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
)

type GoogleProvider struct {
	Verifier *oidc.IDTokenVerifier
}

func NewGoogleProvider() (*GoogleProvider, error) {
	provider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		return nil, fmt.Errorf("NewGoogleProvider err: %w", err)
	}

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	if clientID == "" {
		return nil, errors.New("environment variable GOOGLE_CLIENT_ID is not set")
	}

	oauth2Confing := &oidc.Config{
		ClientID: clientID,
	}

	verifier := provider.Verifier(oauth2Confing)

	return &GoogleProvider{Verifier: verifier}, nil
}
