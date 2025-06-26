package validate

import "errors"

func SendEmailVerificationOtp(email string) (bool, error) {
	if email == "" {
		return false, errors.New("email is required")
	}
	return true, nil
}
