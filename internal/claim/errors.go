package claim

import (
	"errors"
)

var (
	ErrExpiresToken    = errors.New("token is expired")
	ErrInvaidTokenType = errors.New("invalid-token-type")
	ErrInvalidToken    = errors.New("")
)


// this is a coment git 