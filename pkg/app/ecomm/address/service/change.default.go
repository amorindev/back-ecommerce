package service

import (
	"context"
)

func (s *Service) ChangeDefault(ctx context.Context, id string, isDefault bool) error {
	// ? que pasa si no existe el id ver
	_, err := s.AddressRepo.Get(ctx, id)
	if err != nil {
		return err
	}

	return s.AddressRepo.ChangeDefault(ctx, id, isDefault)
}
