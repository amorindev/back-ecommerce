package mongo

import (
	"context"
	"errors"
)

func (r *Repository) RevokedByAuth(ctx context.Context, authID string) error {
	return errors.New("session mongo repo - RevokedByAuth unimplement")
}
