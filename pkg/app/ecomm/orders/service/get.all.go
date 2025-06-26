package service

import (
	"context"

	"com.fernando/pkg/app/ecomm/orders/model"
)

func (s *Service) GetAll(ctx context.Context, userID string) ([]*model.Order, error) {
	// validaciones, paginacion

	orders, err := s.OrderRepo.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	

	return orders, nil
}
