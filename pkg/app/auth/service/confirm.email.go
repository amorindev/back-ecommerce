package service

import "errors"

func (s *Service) ConfirmEmail(userID string) error {
	// Buscar al usuario primero?

	//return s.AuthRepo.ConfirmEmail(userID)
	return errors.New("auth service - ConfirmEmailUnimplement")
}

/*
- deberia hacer una consulta antes o solo ir a actualizar

user, err := s.userRepository.FindByID(userID)
    if err != nil {
        return fmt.Errorf("user not found: %w", err)
    }
    if user.EmailVerified {
        return nil // Email ya está verificado
    }

	- que se devería devolver o nada diseño de la api
    user.EmailVerified = true
    return s.userRepository.Update(user)
*/
