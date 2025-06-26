package port

import (
	"context"

	authModel "com.fernando/pkg/app/auth/model"
		otpM "com.fernando/pkg/app/otp-codes/model"
	userM "com.fernando/pkg/app/user/model"
)

type AuthTransaction interface {
	// este es cuando existe el user y no  el provider password
	// será crear el auth con el otp
	SignUpWithOtp(ctx context.Context, auth *authModel.Auth, otpCode *otpM.OtpCodes) error
	SignUpUser(ctx context.Context, user *userM.User) error
	SignUpWithOtpUser(ctx context.Context, user *userM.User,otpCode *otpM.OtpCodes) error
}



