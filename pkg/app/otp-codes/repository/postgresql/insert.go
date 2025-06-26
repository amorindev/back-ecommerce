package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/otp-codes/model"
)

func (r *Repository) Insert(ctx context.Context, otp *model.OtpCodes) error {
	return errors.New("otp repo - Create is not implement")
}
