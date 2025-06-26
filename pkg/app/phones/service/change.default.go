package service

import "context"

func (s *Service) ChangeDefault(ctx context.Context, id string, isDefault bool) error {
	_, err := s.PhoneRepo.Get(ctx, id)
	if err != nil {
		return err
	}
	return s.PhoneRepo.ChangeDefault(ctx, id, isDefault)
}
