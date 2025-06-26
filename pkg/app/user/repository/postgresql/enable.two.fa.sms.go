package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/user/constants"
)

func (r *Repository) EnableTwoFaSms(ctx context.Context, userID string, twoFaMethod constants.TwoFaMethod) error {
	return errors.New("user pg repo EnableTwoFaSms unimplement")
}