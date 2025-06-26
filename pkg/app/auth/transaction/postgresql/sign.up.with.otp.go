package postgresql

import (
	"context"

	authModel "com.fernando/pkg/app/auth/model"
	"com.fernando/pkg/app/otp-codes/model"
)

func (t *Transaction) SignUpWithOtp(ctx context.Context, auth *authModel.Auth, otpCode *model.OtpCodes) error {
	return nil
}