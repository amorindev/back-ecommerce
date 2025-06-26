package port

import (
	"context"

	"com.fernando/pkg/app/otp-codes/model"
)

type OtpRepo interface {
	Insert(ctx context.Context, otp *model.OtpCodes) error
	Get(ctx context.Context, otpID string) (*model.OtpCodes, error)
	Delete(ctx context.Context, id string) error
}
