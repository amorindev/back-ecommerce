package service

import (
	"context"

	"com.fernando/pkg/app/ecomm/stores/model"
)

func (s *Service) GetAll(ctx context.Context) ([]*model.Store, error) {
	return s.StoreRepo.GetAll(ctx)
}