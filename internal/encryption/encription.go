package encryption

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// refresh token con AES por e número de caracteres
// la clave debe ir en variables de entorno?

// convert password to hash string -- debe ir solo para entidad auth o user si se usa solo para contrseñas

// solo para User -
func HashPassword(password string) (string, error) {
	// admite 72 caracteres probar con el refreshToken
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password %w", err)
	}
	return string(hashPassword), nil
}

// check password is vaid or not, deberíamos agregarle un error con mensaje?
// de esta manera ya no esta acoplado al user
func CheckPassword(password string, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
