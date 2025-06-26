package service

import (
	"context"

	"com.fernando/pkg/app/ecomm/address/model"
)

func (s *Service) GetAll(ctx context.Context, userID string) ([]*model.Address, error) {
	return s.AddressRepo.GetAll(ctx, userID)
}
