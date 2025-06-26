package postgresql

import (
	otpPort "com.fernando/pkg/app/otp-codes/port"
	"github.com/jackc/pgx/v5"
)

var _ otpPort.OtpRepo = &Repository{}

type Repository struct{
	Conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *Repository{
	return &Repository{
		Conn: conn,
	}
}