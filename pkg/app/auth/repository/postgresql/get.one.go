package postgresql

import (
	"context"
	"errors"

	authModel "com.fernando/pkg/app/auth/model"
	//userModel "com.fernando/pkg/app/user/model"
)

func (r *Repository) GetOne(ctx context.Context, id string) (*authModel.Auth, error) {
	/* q := `SELECT id, email, created_at, updated_at FROM tb_auth WHERE id = $1;`

	row := r.Conn.QueryRow(ctx, q, id)

	var a authModel.Auth
	var u userModel.User

	err := row.Scan(&u.ID, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	a.UserAgregate = &u

	return &a, nil */
	return nil, errors.New("auth pg repo - GetOne unimplement")
}
