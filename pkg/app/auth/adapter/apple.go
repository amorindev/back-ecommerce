package adapter

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// https://appleid.apple.com/auth/token
type AppleProvider struct {
	ClientID    string
	TeamID      string
	KeyID       string
	PrivateKey  *rsa.PrivateKey // * validar nil
	AppleKeyUrl string
}

// ! llamar al os.geten desde NewAppleProvider  o desde fuera
// !para wl apple key Url
func NewAppleProvider(clientID, teamID, keyID, privateKey string) (*AppleProvider, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return nil, fmt.Errorf("auth adapter - NewAppleProvider err: %w", err)
	}

	keyURL := os.Getenv("APPLE_KEY_URL")
	if keyURL == "" {
		return nil, errors.New("environment variable APPLE_KEY_URL is not set")
	}

	/* provider, err := oidc.NewProvider(context.Background(), "https://appleid.apple.com/auth/token")
	if err != nil {
		return "", fmt.Errorf("auth adapter - AppleValidateToken err: %w", err)
	}

	oauth2Config := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		// ! ver cmo manejarlo desde flutter
		// !this redirect Url must be set in your Identity Provider Portal
		RedirectURL: "",
		Endpoint:    provider.Endpoint(),
		Scopes:      []string{oidc.ScopeOpenID, "profile", "email"},
	} 
	oidcConfig := &oidc.Config{
		ClientID: "",
	}
	verifier := provider.Verifier(oidcConfig)

	test := oauth2.Endpoint{
		AuthURL: oauth2Config.RedirectURL,
	}	
	*/

	return &AppleProvider{
		ClientID:    clientID,
		TeamID:      teamID,
		KeyID:       keyID,
		PrivateKey:  key,
		AppleKeyUrl: "https://appleid.apple.com/auth/keys",
	}, nil
}

// * como hacer logout ?	
// * un solo hndler pra el sign out con providers, en header con el provider?
/* 
Si la revocación de tokens en los proveedores externos no es obligatoria para ti, un solo handler que elimine la sesión local (invalidate JWT/refresh token) es suficiente. Si necesitas revocar tokens en los proveedores, podrías hacer un sistema modular donde el mismo handler llame a funciones específicas para cada proveedor cuando sea necesario.
*/

/* oicd chatgpt
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/apple"
)

var (
	oauthConfig = oauth2.Config{
		ClientID:     "tu.client.id.de.apple",
		ClientSecret: "tu-client-secret",
		RedirectURL:  "http://localhost:8080/auth/apple/callback",
		Endpoint:     apple.Endpoint,
		Scopes:       []string{"openid", "email"},
	}
	provider *oidc.Provider
)

func appleLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL("state-random")
	http.Redirect(w, r, url, http.StatusFound)
}

func appleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	code := r.URL.Query().Get("code")

	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Token exchange failed", http.StatusInternalServerError)
		return
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "Missing id_token", http.StatusInternalServerError)
		return
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: oauthConfig.ClientID})
	parsedIDToken, err := verifier.Verify(ctx, idToken)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	//fmt.Fprintf(w, "Apple ID: %s", parsedIDToken.Subject)
}

func main() {
	var err error
	provider, err = oidc.NewProvider(context.Background(), "https://appleid.apple.com")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/auth/apple", appleLoginHandler)
	http.HandleFunc("/auth/apple/callback", appleCallbackHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

*/