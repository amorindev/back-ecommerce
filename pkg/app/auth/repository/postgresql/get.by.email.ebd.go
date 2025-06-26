package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/auth/model"
)

func (r *Repository) GetByEmailEbd(ctx context.Context, email string) (*model.Auth, error) {
	return nil, errors.New("auth postgresql repo - GetByEmailEbd unimplement")
}
