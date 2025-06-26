package postgresql

import (
	"context"
	"errors"
)

func (r *Repository) RemoveRolesToUser(ctx context.Context, userID string, roleID string) error {
	return errors.New("role postgresql repo - RemoveRolesToUser unimplement")
}
