package service

import (
	"context"

	"com.fernando/pkg/app/phones/model"
)

func (s *Service) GetAll(ctx context.Context, userID string) ([]*model.Phone, error) {
	return s.PhoneRepo.GetAll(ctx, userID)
}
