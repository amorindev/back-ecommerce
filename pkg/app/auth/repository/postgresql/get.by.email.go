package postgresql

import (
	"context"
	//"database/sql"
	"errors"

	authModel "com.fernando/pkg/app/auth/model"
	/* authErr "com.fernando/pkg/app/auth/errors"
	userModel "com.fernando/pkg/app/user/model" */
)

// ! el unico modelo debe ser Auth

func (r *Repository) GetByEmail(ctx context.Context, email string) (*authModel.Auth, error) {

	/* q := `SELECT a.id, a.provider, a.email, a.password, a.created_at, a.updated_at, u.name, u.username 
			FROM tb_auth a 
			FULL JOIN tb_user u ON a.id = u.id
			WHERE a.email = $1;
		`

	row := r.Conn.QueryRow(ctx, q, email)

	var a authModel.Auth
	var u userModel.User

	//! no hay name ni username
	err := row.Scan(&u.ID, &a.Provider, &u.Email, &a.PasswordHash, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, authErr.ErrAuthNotFound
		}
		// interval server error? = error al obtener el usuario
		return nil, fmt.Errorf("get user postgresql err: %w", err)

	}

	a.UserAgregate = &u

	return &a, nil */
	return nil, errors.New("auth pg repo - GetByEmail unimplement")
}

// ! refactorizar por qu los get tran el auth no agregan al auth es asi como lo estamos manejando
// auth debe tener el email
func (r Repository) GetByEmail2(ctx context.Context, auth *authModel.Auth) error {
	/* q := `SELECT a.id, a.email, a.password, a.created_at, a.updated_at 
			FROM tb_auth a WHERE email = $1
			INNER JOIN tb_user u ON a.id = u.id;
		`

	row := r.Conn.QueryRow(ctx, q, auth.UserAgregate.Email)

	// !verificar esta funcion a.User = &user
	//var a authModel.Auth
	// ! es la responsabilidad de UserRepo
	var a authModel.Auth
	var u userModel.User

	err := row.Scan(&u.ID, &u.Email, &a.PasswordHash, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			auth = nil

			return errors.New("user-not-found")
		}
		// interval server error? = error al obtener el usuario
		auth = nil
		return fmt.Errorf("get user postgresql err: %w", err)

	}
	// ? a = &u y luego auth = a
	auth.UserAgregate = &u
	return nil */
	return errors.New("auth pg repo - GetByEmail2 unimplement")

}
