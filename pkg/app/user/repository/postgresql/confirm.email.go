package postgresql

import "errors"

// !esto es mal por que es de la tabla user y ademas se debe hacer por userID o email
// ? deberia retornar bool? --
// ? esta funcion no se parece el update user o auth, tambien se debe asegurar no mandandar campos
// que no se deben actuaizar
func (r *Repository) ConfirmEmail(email string) error {
	/* q := `UPDATE tb_auth SET email_verified = true WHERE id = $1 AND email_verified = false`

	_, err := r.Conn.Exec(context.Background(), q, authID)
	if err != nil {
		if err == sql.ErrNoRows {
			// se devolvería al frontend?
			return fmt.Errorf("confirmEmmail - AuthRepository: User not found or user confirmed")
		}
		return fmt.Errorf("confirmEmmail - AuthRepository: %w", err)
	}
	*/
	return errors.New("user postgresql repo - ConfirmEmail unimplement")
}
