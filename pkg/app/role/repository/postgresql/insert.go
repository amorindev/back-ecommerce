package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/role/model"
)

func (r *Repository) Insert(ctx context.Context, role *model.Role) error {
	return errors.New("role postgresql repo - Create unimplement")
}
