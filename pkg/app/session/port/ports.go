package port

import (
	"context"

	"com.fernando/pkg/app/session/model"
)

// cambiar a sesiones

type SessionRepo interface {
	Create(ctx context.Context, session *model.Session) error
	Delete(ctx context.Context, id string) error
	// si usas el jtti del token lo tienes que buscar por RefreshTokenID
	Get(ctx context.Context, id string) (*model.Session, error)
	// * Recordemos que debemos buscarlo por refreshtokenid
	GetByClaimID(ctx context.Context, id string) (*model.Session, error)
	// ! eliminar por que se relaciona  con user
	GetByAuth(ctx context.Context, authID string) ([]model.Session, error)
	RevokedByAuth(ctx context.Context, authID string) error
}

type SessionSrv interface {
	Create(session *model.Session, roles []string) error
	// * Recordemos que no usamos el id principal sino el campo refreshtokenid
	// * esta dentro de jwt.RegisterClaims en su propiedad ID
	RefreshToken(ctx context.Context, refreshTokenID string) (string, string, error)
}
