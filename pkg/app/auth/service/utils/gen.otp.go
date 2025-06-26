package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// crear el token para send email verification
// deberia estar en internal= utils?
func GenOtpCode() (string, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(1000000)) // 0 to 999999 inclusive
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", nBig.Int64()), nil
}
