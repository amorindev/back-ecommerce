package service

import "context"

func (s *Service) RevokeToken(ctx context.Context, sessionID string) error {

	// ? se debe manejar el error ? o el cron job se encargará de eliminarlo
	go s.SessionRepo.Delete(context.Background(), sessionID)

	return nil
}
