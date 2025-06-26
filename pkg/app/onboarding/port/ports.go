package port

import (
	"context"

	"com.fernando/pkg/app/onboarding/model"
)

// * Crearlo desde el init
type OnboardingRepo interface {
	Insert(ctx context.Context, onboarding *model.Onboarding) error
	Get(ctx context.Context) ([]*model.Onboarding, error)
	// esto sera el correcto ver si va a usar paginacion
	//GetAll(ctx context.Context) ([]*model.Onboarding, error)
	GetByTitle(ctx context.Context, title string) (*model.Onboarding, error)
}

// *Faltaría marcar como visto
type OnboardingSrv interface {
	Create(ctx context.Context, onboarding *model.Onboarding) error
	// get para obtener todos y getall con paginacion Get para pruebas
	Get(ctx context.Context) ([]*model.Onboarding, error)
	// ? paginacion? solo 3
	//GetAll(ctx context.Context) ([]*model.Onboarding, error)
}


