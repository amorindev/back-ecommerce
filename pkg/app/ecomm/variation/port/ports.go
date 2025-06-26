package port

import (
	"context"

	"com.fernando/pkg/app/ecomm/variation/model"
)

type VariationRepo interface {
	Create(ctx context.Context, variation *model.Variation) error
	CreateMany(ctx context.Context, variations []*model.Variation) error
	GetByName(ctx context.Context, name string) (*model.Variation, error)
}

type VariationService interface {
	GetWithOptions(ctx context.Context) error
}
