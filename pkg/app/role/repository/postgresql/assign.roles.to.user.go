package postgresql

import (
	"context"
	"errors"
	"fmt"

	"com.fernando/pkg/app/role/model"
	"github.com/google/uuid"
)

// * Begin o BeginTx repasar, libpq no piene preatare pero pgx si, revisar quien no tenía
// revisa tx.
// * como usar los context con pgx y la capa de trasaccion
func (r *Repository) AssignRolesToUser(ctx context.Context, userID string, roles []model.Role) error {
	tx, err := r.Conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("role postgresql repo - AssignRolesToUser err: %w", err)
	}

	// ! esto no era cuando ocurre un error
	defer tx.Rollback(ctx)

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("role postgresql repo - AssignRolesToUser err: %w", err)
	}

	for _, role := range roles {
		roleUUID, err := uuid.Parse(role.ID.(string))
		if err != nil {
			// Rollback?
			return fmt.Errorf("role postgresql repo - AssignRolesToUser err: %w", err)
		}
		_, err = tx.Exec(ctx, "INSERT INTO tb_user_role (user_id, role_id) VALUES ($1,$2)", userUUID, roleUUID)
		if err != nil {
			//Rollback?
			return fmt.Errorf("role postgresql repo - AssignRolesToUser err: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("role postgresql repo - AssignRolesToUser err: %w", err)
	}

	return errors.New("role postgresql repo - review")
}
