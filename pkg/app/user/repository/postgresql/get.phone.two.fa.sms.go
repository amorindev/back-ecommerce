package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/phones/model"
)
func (r *Repository) GetPhonetwoFaSms(ctx context.Context, userID string) (*model.Phone, error) {
	return nil, errors.New("user pg repo - getPhoneTwoFaSms unimplement")
}
