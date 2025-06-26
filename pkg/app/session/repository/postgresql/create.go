package postgresql

import (
	"context"

	"com.fernando/pkg/app/session/model"
	"github.com/google/uuid"
)

// Guardaremos como string en ambos mongo y postgresql
// add device
func (r *Repository) Create(ctx context.Context, session *model.Session) error {
	// deberia pasarlo a string?
	// create id - importante
	// revisar que todos creen de esta manera y no en una variable aparte, (hay ecepciones)
	session.ID = uuid.New()

	// * falta created revisar todos los campos, necesita updated?

	// agregue refresh_token_id a la tabla
	q := `INSERT INTO tb_session (id, refresh_token_id ,refresh_token, expires_at, created_at, revoked, user_id)
		  VALUES($1,$2,$3,$4,$5,$6)
	`
	_, err := r.Conn.Exec(context.Background(), q, session.ID, session.RefreshTokenID, session.RefreshTokenHash, session.ExpiresAt, session.CreatedAt, session.Revoked, session.UserID)
	if err != nil {
		return err
	}

	return nil
}
