package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/session/model"
)

func (r *Repository) GetByClaimID(ctx context.Context, id string) (*model.Session, error) {
	return nil, errors.New("session postgresql repo - GetByClaimID unimplement")
}