package service

import (
	"context"
	"errors"

	"com.fernando/internal/encryption"
	authErr "com.fernando/pkg/app/auth/errors"
	userErr "com.fernando/pkg/app/user/errors"
)

// ! en todo lo que es authenticacion reidar el limit rating tanto en la ui como
// ! en el backend
// * de momemento solo eliminaré la tabla documento de la colleccion user
// * otra opcion escambiar el estado de usuario a adesactivado
func (s *Service) DeleteAccount(ctx context.Context, userID string, password string) error {
	// no seria necesario, ademas el delete tambien retorna que no existe el user
	// ver responsabilidades
	/* user, err := s.UserRepo.Get(ctx, userID)
	if err != nil {
	  return err
	} */

	auth, err := s.AuthRepo.GetByIDProvider(ctx, userID, "password")
	if err != nil {
		if err == authErr.ErrAuthNotFound {
			return userErr.ErrUserNotFound
		}
		return err
	}
	// donde verificar el nil en capa de repo o service
	err = encryption.CheckPassword(password, *auth.PasswordHash)
	if err != nil {
		// ! desde donde regular los mensajes de erro por ejempl
		// ! cuando se hace login no se pone invalid password sino general invalid-credentials
		// ! me parece que lo mas seguro es ya parsearlo y no enviar password do not march
		// ! ver ver el punto adecuado
		// return errors.New("passwords-do-not-match") mejor esta el de abajo
		return errors.New("incorrect-password")
	}

	// que pasa con los proveedores como apple google, cual seria el flujo porque no tiene password
	// otp? no lo se, se tine que hacer algo o eviar algo al proveedor antes de eliminar la cuenta?
	// como facebook no tiene sign in con proveedores usa el password y es mas sencillo
	// Delete te retorna el user-not-found ver el tema de responasbilidades
	err = s.UserRepo.Delete(ctx, userID)
	if err != nil {
		return err
	}

	// me parece mejor adjuntar el ID de la session para despues eliminarla ela hacer logout
	// o delete account dentro del token y con una go rutina eliminarlo

	return nil
}
