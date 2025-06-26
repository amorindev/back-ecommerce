package mongo

import (
	"context"
	"errors"
)

func (r Repository) FindWithRoles(ctx context.Context, authID string) ([]string, error) {
	return nil, errors.New("auth mongo repo - FindWithRoles, unimpement")
}
