package port

import (
	"context"
)

type InitService interface {
	CreateRoles(ctx context.Context, roleNames []string) error
	CreatePermissions() error
	CreateAdmin(ctx context.Context, email string, password string, roleIDs []string) error
	CreateProducts() error
	CreateCategories() error
	CreateVariations() error
	CreateVariationOptions() error
	CreateOnboarding() error
	CreateStores() error
}

// * si hay alguna funcionalidad solo para init
type InitTransaction interface {
}

// * si hay alguna funcionalidad solo para init
type InitRepository interface {
}
