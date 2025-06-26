package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/auth/model"
)

func (r *Repository) CreateWithUser(ctx context.Context, auth *model.Auth) error {
	return errors.New("auth postgresql repo - CreateWithUser unimplement")
}
