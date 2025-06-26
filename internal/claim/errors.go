package claim

import (
	"errors"
)

var (
	ErrExpiresToken    = errors.New("token is expired")
	ErrInvaidTokenType = errors.New("invalid-token-type")
	ErrInvalidToken    = errors.New("")
)
