package postgresql

import (
	"context"

	"com.fernando/pkg/app/session/model"
)

func (r *Repository) GetByAuth(ctx context.Context, authID string) ([]model.Session, error) {
	return nil, nil
}
