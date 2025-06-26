package postgresql

import (
	"context"
	"errors"

	authModel "com.fernando/pkg/app/auth/model"
	//userModel "com.fernando/pkg/app/user/model"
)

func (r *Repository) GetAll(ctx context.Context) ([]authModel.Auth, error) {
	/* q := `SELECT id, email, created_at, updated_at FROM tb_auth;`

	rows, err := r.Conn.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []authModel.Auth

	for rows.Next() {
		var a authModel.Auth
		// ! es la responsabilidad de Userrepo
		var u userModel.User
		err = rows.Scan(&u.ID, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		a.UserAgregate = &u
		users = append(users, a)
	}

	return users, nil */
	return nil, errors.New("auth pg repo - getAll unimplement")
}
