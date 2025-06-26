package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/role/model"
)

func (r *Repository) GetByUserIDRole(ctx context.Context, userID string) ([]model.Role, error) {
	return nil, errors.New("role postgresql repo - GetByUserIDRole unimplement")
}


