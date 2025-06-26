package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/role/model"
)

func (r *Repository) GetByNames(ctx context.Context, names []string) ([]model.Role, error) {
	return nil, errors.New("role postgresql repo - GetByNames unimplement")
}
