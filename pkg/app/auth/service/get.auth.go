package service

import (
	"context"

	"com.fernando/pkg/app/auth/model"
)

func (s *Service) GetAuth(ctx context.Context, authID string) (*model.Auth, error) {
	// si tendríamos el campo de usario activo lo filtrariamos tambien
	// afectaría a otros servicios
	return s.AuthRepo.Get(ctx, authID)
}
