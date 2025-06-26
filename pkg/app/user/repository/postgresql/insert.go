package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/user/model"
)

func (r *Repository) Insert(ctx context.Context, user *model.User) error {
	return errors.New("user postgresql repo - Create unimplement")
}
