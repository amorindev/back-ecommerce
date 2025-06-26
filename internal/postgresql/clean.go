package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// * limpiar la base de datos

func cleanDatabase(conn *pgx.Conn) error {
	return conn.Close(context.Background()); //para cerrar la coneccion desde el server o donde?
}