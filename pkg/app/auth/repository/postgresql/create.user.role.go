package postgresql

import "context"

func (r *Repository) CreateUserRole(ctx context.Context, authID string, roleID string) error {
	//verificar
	q := `INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2)`

	_ = r.Conn.QueryRow(ctx, q, authID, roleID)

	return nil

}
