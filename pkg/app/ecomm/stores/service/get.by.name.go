package service

import (
	"context"

	"com.fernando/pkg/app/ecomm/stores/model"
)

func (s *Service) GetByName(ctx context.Context, name string) (*model.Store, error) {
	return s.StoreRepo.GetByName(ctx, name);
}