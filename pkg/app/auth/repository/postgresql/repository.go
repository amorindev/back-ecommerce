package postgresql

import (
	"com.fernando/pkg/app/auth/port"
	"github.com/jackc/pgx/v5"
)

var _ port.AuthRepo = &Repository{}

// no seria mejor Postgresql Repositoru
type Repository struct {
	Conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *Repository{
	return &Repository{
		Conn: conn,
	}
}