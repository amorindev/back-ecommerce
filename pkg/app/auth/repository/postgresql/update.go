package postgresql

import (
	"context"

	"com.fernando/pkg/app/auth/model"
)

func (r *Repository) Update(ctx context.Context, id string, user model.Auth) error {
	return nil
}