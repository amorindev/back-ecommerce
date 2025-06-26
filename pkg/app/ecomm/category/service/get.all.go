package service

import (
	"context"

	"com.fernando/pkg/app/ecomm/category/model"
)

func (s *Service) GetAll(ctx context.Context) ([]*model.Category, error) {
	categories, err := s.CategoryRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
