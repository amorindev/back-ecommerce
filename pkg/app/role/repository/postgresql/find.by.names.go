package postgresql

import (
	"context"
	"database/sql"
)

// retorna los IDs
// ? deberia parsearlo a string el ID
func (r *Repository) FindByNames(ctx context.Context, names []string) (roleIDs []string, err error) {
	q := `SELECT id FROM tb_role WHERE name = $1;`

	var idRoles []string

	for _, name := range names {
		var id string
		row := r.Conn.QueryRow(ctx, q, name)

		err := row.Scan(&id)
		if err == sql.ErrNoRows {
			//fmt.Println("No rows were returned!")
		}
		if err != nil {
			return nil, err
		}
		idRoles = append(idRoles, id)
	}

	return idRoles, nil
}
