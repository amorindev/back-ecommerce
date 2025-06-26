package postgresql

import (
	"context"
	"errors"
)

func (r *Repository) Delete(ctx context.Context, id string) error {
	return errors.New("otp postgresql repository - Delete not implement")
}
