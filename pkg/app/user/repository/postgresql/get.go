package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/user/model"
)

func (r *Repository) Get(ctx context.Context, id string) (*model.User, error) {
	return nil, errors.New("user postgresql repo - Get unimplement")
}
