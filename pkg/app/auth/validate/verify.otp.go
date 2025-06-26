package validate

import "errors"

// ? variable de entorno para cantidad de digitos en OTP?
// * hay un error generico para todos los campos verificar cuadros
func ValdiateVerifyOtp(otpID, otpCode, userID, email string) (bool, error) {
	if otpID == "" {
		return false, errors.New("otp_id is required")
	}

	if otpCode == "" {
		return false, errors.New("otp_code is required")
	}
	/* if userID == "" {
		return false, errors.New("user_id is required")
	} */
	if email == "" {
		return false, errors.New("email is required")
	}

	return true, nil
}
