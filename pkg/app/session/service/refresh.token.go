package service

import (
	"context"
	"errors"

	"com.fernando/pkg/app/session/model"
)

// ! desde el middleware  no desde los servicios claim.GetConfig() buscar vscode
// ! identificar en que servicios se va crear os tokens
// * de momento devolveremos string  string
// * puedo devolver el Session{} pero estarimos acoplandolo al auth dependencia
// * cuales son los puntos que se deja pasar y cuales el usario deberá hacer sign in nuevamente
func (s *Service) RefreshToken(ctx context.Context, refreshTokenID string) (string, string, error) {

	// * verificar si existe el refresh en la base de datos
	session, err := s.SessionRepository.GetByClaimID(ctx, refreshTokenID)
	if err != nil {
		return "", "", err
		// si no existe el errors es session-revoked
	}

	// ? esto no lo valida el middleware
	if session.ExpiresAt.After(session.ExpiresAt) {
		return "", "", errors.New("session-expired-refresh-token")
	}

	// ! cual usar session o  refresh_tokens coleccion
	if session.Revoked {
		return "", "", errors.New("session-revoked")
	}

	// Rotar el refresh token - rotar signica traer generar uno nuevo y guardarlo en la base dedatos
	// bueno eliminar podría estar cerca del concepto
	// * go para los de redis y tambien el cronjob

	go s.SessionRepository.Delete(context.Background(), session.ID.(string))

	// no se necesitaria esta funcion si el refreshtoken tendria el email es inseguro?
	// o sacrificar labase de datos

	user, err := s.UserRepository.Get(ctx, session.UserID.(string))
	if err != nil {
		return "", "", err
	}

	roles, err := s.RoleRepository.GetByUserID(ctx, user.ID.(string))
	if err != nil {
		return "", "", err
	}

	newSession := &model.Session{
		UserID:    session.UserID.(string),
		Email:     user.Email,
		// * De momento sin refresh
		RemenberMe: false,
	}

	err = s.Create(newSession,roles)
	if err != nil {
		return "", "", err
	}

	// en estepunto Create de session service ya lo asigno
	return newSession.AccessToken, newSession.RefreshToken, nil
}

/*
var userID string
	var expiresAt time.Time
	var revoked bool

	err := db.QueryRow(`
        SELECT user_id, expires_at, revoked
        FROM refresh_tokens
        WHERE token = $1`, token).Scan(&userID, &expiresAt, &revoked)
if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("token no encontrado")
		}
		return "", err
	}
*/

//Cuando el cliente envía un refresh token, validas:
//Que el token exista y no haya expirado.
//Que no haya sido revocado.

//Si es válido, generas un nuevo access token y, si es necesario, un nuevo refresh token.
