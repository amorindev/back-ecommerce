package postgresql

import (
	"context"
	"database/sql"

	"com.fernando/pkg/app/role/model"
)

func (r *Repository) FindByNamesRole(ctx context.Context, names []string) ([]model.Role, error) {
	q := `SELECT id, name FROM tb_role WHERE name = $1;`

	var roles []model.Role

	for _, name := range names {
		var role model.Role
		row := r.Conn.QueryRow(ctx, q, name)

		err := row.Scan(&role.ID, &role.Name)
		if err == sql.ErrNoRows {
			//fmt.Println("No rows were returned!")
		}
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}
