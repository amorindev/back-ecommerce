package handler

import "net/http"
func (h Handler) AppleSignIn(w http.ResponseWriter, r *http.Request) {
	/*
	claims, err := s.appleProvider.ValidateToken(appleToken)
	if err != nil {
		return "", err
	}

		// Buscar usuario en la base de datos
	user, _ := s.userRepo.FindByEmail(ctx, claims.Email)
	if user == nil {
		// Crear nuevo usuario y auth si no existe
		newUser := user.User{
			ID:            generateUUID(),
			Name:          claims.Name,
			Username:      generateUsername(claims.Email),
			Email:         claims.Email,
			EmailVerified: true,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		newAuth := auth.Auth{
			ID:        generateUUID(),
			UserID:    newUser.ID,
			Provider:  "apple",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := s.createUserWithAuth(ctx, newUser, newAuth)
		if err != nil {
			return "", err
		}

		user = &newUser
	*/
	/*generar el token*/
}

/*
chat


func VerifyAppleToken(idToken string) (map[string]interface{}, error) {
	// Apple uses public keys for token validation.
	resp, err := http.Get("https://appleid.apple.com/auth/keys")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Apple public keys: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Apple public keys: %v", err)
	}

	var keys map[string]interface{}
	if err := json.Unmarshal(body, &keys); err != nil {
		return nil, fmt.Errorf("failed to parse Apple keys: %v", err)
	}

	// Validate token here (parsing, matching public keys, etc.)
	// For simplicity, this assumes a library is used to validate JWTs.
	return nil, nil
}
*/
