package postgresql

import (
	"context"

	roleErr "com.fernando/pkg/app/role/errors"
	"com.fernando/pkg/app/role/model"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetByName(ctx context.Context, name string) (*model.Role, error) {
	q := `SELECT id FROM tb_role WHERE name = $1;`

	row := r.Conn.QueryRow(ctx, q, name)

	var role model.Role

	err := row.Scan(&role.ID, &role.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, roleErr.ErrRoleNotFound
		}
		return nil, err
	}

	return &role, nil
	
	//return nil, errors.New("role postgresql repo - GetByName unimplement")
}
