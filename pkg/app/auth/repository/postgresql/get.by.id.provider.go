package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/auth/model"
)

func (r *Repository) GetByIDProvider(ctx context.Context, userID string, provider string) (*model.Auth, error) {
	return nil, errors.New("auth mongo repo - GetByIDProvider unimplement")
}
