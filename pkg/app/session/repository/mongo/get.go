package mongo

import (
	"context"
	"errors"

	"com.fernando/pkg/app/session/model"
)

func (r *Repository) Get(ctx context.Context, id string) (*model.Session, error) {
	return nil, errors.New("session mongo repo - Get unimplement")
}
