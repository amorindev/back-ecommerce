package service

func (s *Service) AppleSignIn() error {
	// El cliente obtiene un token JWT desde google/apple
	// El backend verifica este token con Google/Apple
	// Si el token es válido, busca al usuario en la tabla auth o lo registra si es un nuevo usuario
	// Genera un token jwt propio para e usaurio y o guarda en la tabla session.

	return nil
}

// ver si verify email en auth changes ver supabase sdk
