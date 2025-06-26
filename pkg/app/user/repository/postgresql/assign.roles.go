package postgresql

import (
	"context"
	"errors"
)

func (r *Repository) AssignRoles(ctx context.Context, userID string, roleIDs []string) error {
	return errors.New("user postgresql repo - AssignRoles unimplement")
}
