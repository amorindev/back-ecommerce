package postgresql

import (
	"context"
	"errors"

	otpM "com.fernando/pkg/app/otp-codes/model"
	userM "com.fernando/pkg/app/user/model"
)

func (t *Transaction) SignUpWithOtpUser(ctx context.Context, user *userM.User, otpCode *otpM.OtpCodes) error {
	return errors.New("auth pg repo - SignUpWithOtpUser unimplement")
}
