package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/otp-codes/model"
)

func (r *Repository) Get(ctx context.Context, otpID string) (*model.OtpCodes, error) {
	return nil, errors.New("otp postgresql repository - Get not implement")
}
