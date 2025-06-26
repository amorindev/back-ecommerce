package postgresql

import (
	"context"
	"fmt"
)

func (r *Repository) GetByUserID(ctx context.Context, userID string) ([]string, error) {
	/* q := `
		SELECT r.name FROM tb_role r
		JOIN tb_auth_role ur ON r.id = ur.id
		JOIN tb_auth u ON ur.id = u.id
		WHERE s.id = $1;
	` */

	q := `SELECT r.name 
			FROM tb_role r
			JOIN tb_auth_role ar ON r.id = ar.role_id
			JOIN tb_auth a ON ar.auth_id = a.id
			WHERE a.id = $1;`

	rows, err := r.Conn.Query(ctx, q, userID)
	if err != nil {
		return nil, fmt.Errorf("GetByUserID err: %w", err)
	}

	defer rows.Close()

	var roles []string

	for rows.Next() {
		var r string
		err := rows.Scan(&r)
		if err != nil {
			return nil, fmt.Errorf("GetByUserID scan error: %w", err)
		}
		roles = append(roles, r)
	}

	return roles, nil
}
