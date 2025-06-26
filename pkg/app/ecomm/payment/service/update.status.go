package service

import (
	"context"

	"com.fernando/pkg/app/ecomm/payment/model"
)

func (s *Service) UpdateStatus(ctx context.Context, id string, status model.PaymentStatus) error {
	return s.PaymentRepo.UpdateStatus(ctx, id,status)
}