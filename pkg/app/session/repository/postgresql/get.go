package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"com.fernando/pkg/app/session/model"
)

// que datos deberiamos traer todos? - sin el refresh token igual a getOne del user sin la contraseña
func (r *Repository) Get(ctx context.Context, id string) (*model.Session, error) {
	// get by id de lata tabla y el id del refresh token

	q := `SELECT id, refresh_token, expires_at, created_at, revoked, user_id FROM tb_session WHERE id = $1;
	`

	row := r.Conn.QueryRow(ctx, q, id)

	var t model.Session

	err := row.Scan(&t.ID, &t.RefreshTokenHash, &t.ExpiresAt, &t.CreatedAt, &t.Revoked, &t.UserID)
	if err != nil {
		// validar si esto funciona
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("token not found, %w", err)
		}
		return nil, fmt.Errorf("get token postgresql error: %w", err)
	}

	/* fmt.Printf("Token ID: %v\n", t.ID)
	fmt.Printf("Token RefreshTokenHash: %v\n", t.RefreshTokenHash)
	fmt.Printf("Token ExpiresAt: %v\n", t.ExpiresAt)
	fmt.Printf("Token CreatedAt: %v\n", t.CreatedAt)
	fmt.Printf("Token Revoked: %v\n", t.Revoked)
	fmt.Printf("Token userID: %v\n", t.UserID) */

	return &t, nil
}
