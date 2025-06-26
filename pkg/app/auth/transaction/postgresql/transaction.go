package postgresql

import (
	authPort "com.fernando/pkg/app/auth/port"
	rolePort "com.fernando/pkg/app/role/port"
	"github.com/jackc/pgx/v5"
)

var _ authPort.AuthTransaction = &Transaction{}

type Transaction struct {
	Client   *pgx.Conn
	AuthRepo authPort.AuthRepo
	RoleRepo rolePort.RoleRepo
}

func NewTransaction(client *pgx.Conn, authRepo authPort.AuthRepo, roleRepo rolePort.RoleRepo) *Transaction {
	return &Transaction{
		Client:   client,
		AuthRepo: authRepo,
		RoleRepo: roleRepo,
	}
}
