package postgresql

import (
	"com.fernando/pkg/app/session/port"
	"github.com/jackc/pgx/v5"
)

var _ port.SessionRepo = &Repository{}

type Repository struct {
	Conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{
		Conn: conn,
	}
}
