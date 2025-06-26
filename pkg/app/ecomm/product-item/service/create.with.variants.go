package service

import (
	"context"
	"errors"

	"com.fernando/pkg/app/ecomm/products/model"
)

func (s *Service) CreateWithVariants(ctx context.Context, product *model.Product) error {
	return errors.New("product mongo repo - CreateWithVariants unimplement")
}
