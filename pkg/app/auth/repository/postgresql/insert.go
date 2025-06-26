package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/auth/model"
	//"github.com/google/uuid"
)

// * transccion aqui? solo debería insertar auth, y como hacer con el embedding?
// TODO: Hay prepare en pgx
// ? pgx utiliza el preparado por defecto
// ? cuando usar Err sql no rows
func (r *Repository) Create2(ctx context.Context, a *model.Auth) error {
	/* authID := uuid.New()
	a.UserAgregate.ID = authID.String()
	// version corta ver el tema siguiente
	// ? se debe insertar el uuid tipo  o el string?
	// probar con ambos
	//a.User.ID = uuid.New().String()

	// * Insert user data

	tx, err := r.Conn.Begin(context.Background())
	if err != nil {
		// go senior asi para todos los errores
		tx.Rollback(context.Background())
		return fmt.Errorf("Create repo - error al iniciar la transacción, %w", err)
	}

	// ! ver go senior si lo usa analizar
	//defer tx.Rollback(context.Background())

	// ! no hay name ni username
	q := `
		INSERT INTO tb_user (id, name, username)
		VALUES ($1, $2, $3);
	`

	_, err = r.Conn.Exec(ctx, q, a.UserAgregate.ID)
	if err != nil {
		return err
	}

	q = `
		INSERT INTO tb_auth (id, email, password, email_verified, provider,created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	_, err = r.Conn.Exec(ctx, q, a.UserAgregate.ID, a.UserAgregate.Email, &a.PasswordHash, a.UserAgregate.EmailVerified, &a.UserAgregate.CreatedAt, &a.UserAgregate.CreatedAt, &a.UserAgregate.UpdatedAt)
	if err != nil {
		return err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("auth mongo repo - Create 1 or 2: %w", err)
	}

	return nil */
	return errors.New("auth pg repo - Create2 unimplement")
}

func (r *Repository) Insert(ctx context.Context, a *model.Auth) error {
	/* id := uuid.New()
	// ? o sin .String()?
	a.ID = id.String()
	a.UserAgregate.ID = id.String()

	// * Insert user data

	tx, err := r.Conn.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("Create repo - error al iniciar la transacción, %w", err)
	}

	defer tx.Rollback(context.Background())

	q := `
		INSERT INTO tb_user (id, name, username)
		VALUES ($1, $2, $3);
	`

	// ! no hay name ni username
	_, err = tx.Exec(ctx, q, a.UserAgregate.ID)
	if err != nil {
		return fmt.Errorf("Create repo - error insert user, %w", err)
	}

	// ? Si es sign in con google como obtengo sus datos?
	q = `
		INSERT INTO tb_auth (id, email, password, email_verified, provider,created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	_, err = tx.Exec(ctx, q, a.UserAgregate.ID, a.UserAgregate.Email, a.PasswordHash, a.UserAgregate.EmailVerified, &a.Provider, &a.UserAgregate.CreatedAt, &a.UserAgregate.UpdatedAt)

	if err != nil {
		return fmt.Errorf("Create repo - error insert auth, %w", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("Create repo - error commit, %w", err)
	}

	return nil */
	return errors.New("auth pg repo - Insert unimplment")
}

/*

func (ur *UserRepository) Create(ctx context.Context, u *userentities.User) error {
	userID := uuid.New()
	u.ID = userID.String()

	q := `
		INSERT INTO tb_auth (id, email, password, is_verified, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5,$6)
		RETURNING id;
	`

	// pgx utiliza el preparado por defecto
	_ = ur.Conn.QueryRow(ctx, q, u.ID, u.Email, u.PasswordHash, u.EmailVerified, u.CreatedAt, u.UpdatedAt)

	return nil
}

*/
