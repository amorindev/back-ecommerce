package postgresql

import (
	"context"
	"errors"
)

// deberia estar en user DDD
func (r *Repository) Delete(ctx context.Context, id string) error {
	return errors.New("user postgresql repo - Delete unimplement")
}
