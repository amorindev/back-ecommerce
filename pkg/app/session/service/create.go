package service

import (
	"context"
	"errors"
	"time"

	"com.fernando/internal/claim"
	"com.fernando/pkg/app/session/model"
)

// usar generateJTI o uuid pero genera dependencia
// esta funcion donde debe vivir?
// como serifico si es único? mejor uuid
/* func generateJTI() (string, error) {
	return "", nil
} */

// * estoy usando session en
// sign-in sign-up refresh-token

// * Donde se necesita
// userID
// UserRoles
// RememberMe
// Email revisar que todos los que usan lo tengan

// * Session service
// CreatedAt
//

// * Database
// ID
// RefreshTokenID

// ! eliminarrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr
// ! eliminarrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr
// ! eliminarrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr
// ! eliminarrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr
// ! comentalo y vas viendo donde se malogra

// ! eliminar y mantener un solo servicio despues lo desacoplas
// ver como sacar el expires in del access token o refresjtoken?
func (s *Service) Create(session *model.Session, roles []string) error {
	now := time.Now().UTC()

	session.CreatedAt = &now

	err := session.HashToken()
	if err != nil {
		return err
	}

	conf, err := claim.GetConfig()
	if err != nil {
		// internal server error?
		return errors.New("internal sginingString or issuer not set")
	}

	// * AccessToken
	accessClaim := claim.NewAccessToken(session.UserID.(string), session.Email, conf.Issuer, roles, conf.Audience)

	accessToken, err := accessClaim.GetToken(conf.AccessString)
	if err != nil {
		return err
	}

	// * RefreshToken
	refreshClaim := claim.NewRefreshToken(session.UserID.(string), session.RemenberMe)

	refreshToken, err := refreshClaim.GetToken(conf.RefreshString)
	if err != nil {
		return err
	}

	session.RefreshTokenID = refreshClaim.ID
	session.RefreshToken = refreshToken
	session.AccessToken = accessToken
	session.ExpiresAt = refreshClaim.ExpiresAt.Time
	session.ExpiresIn = int64(time.Until(accessClaim.ExpiresAt.Time).Seconds())
	session.Revoked = false // ! ver los demas por que se insertarán dependienndo nil o ""

	err = s.SessionRepository.Create(context.Background(), session)
	if err != nil {
		return err
	}

	return nil
}
