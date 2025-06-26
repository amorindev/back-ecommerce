package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/user/model"
)

func (r *Repository) CreateEbd(ctx context.Context, user *model.User) error {
	return errors.New("user postgresql repo - CreateEbd unimplement")
}
