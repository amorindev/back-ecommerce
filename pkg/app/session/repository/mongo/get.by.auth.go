package mongo

import (
	"context"
	"errors"

	"com.fernando/pkg/app/session/model"
)

func (r *Repository) GetByAuth(ctx context.Context, authID string) ([]model.Session, error) {
	return nil, errors.New("session mongo repo - GetByAuth unimplement")
}
