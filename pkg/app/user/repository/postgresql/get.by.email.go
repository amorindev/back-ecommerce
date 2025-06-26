package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/user/model"
)

func (r *Repository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, errors.New("user postgresql repo - GetByEmail unimplement")

}
