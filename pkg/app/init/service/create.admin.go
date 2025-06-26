package service

import (
	"context"
	"time"

	authModel "com.fernando/pkg/app/auth/model"
	userErr "com.fernando/pkg/app/user/errors"
	userM "com.fernando/pkg/app/user/model"
)

// /los datos delUser?
func (s *Service) CreateAdmin(ctx context.Context, email string, password string, roleNames []string) error {
	now := time.Now().UTC()
	provider := "password"

	// ? verificar si ya existe el usuario o una lista de emails? seria seguro?

	// ! falta name y username si no los agrego que pasa por que no son punteros
	// ! dos cosas o sera "" o error
	userCreate := &userM.User{
		Email:         email,
		EmailVerified: false,
		CreatedAt:     &now,
		UpdatedAt:     &now,
		/* AuthProviders: []*authModel.Auth{
			{
				Provider:  &provider,
				Password:  password,
				CreatedAt: &now,
				UpdatedAt: &now,
			},
		}, */
		AuthProviderCreate: &authModel.Auth{
			Provider:  &provider,
			Password:  password,
			CreatedAt: &now,
			UpdatedAt: &now,
		},
	}
	//auth.UserAgregate = user
	//err := userCreate.AuthProviders[0].HashPassword()
	err := userCreate.AuthProviderCreate.HashPassword()
	if err != nil {
		return err
	}
	// aqui no es como el sign up por que el usario no debe existir
	user, err := s.UserRepo.GetByEmail(ctx, email)
	if err != nil {
		if err != userErr.ErrUserNotFound {
			return err
		}
	}

	if user != nil {
		//return fmt.Errorf("user already exists %s", user.Email)
		return nil
	}

	// * Assign roles

	// desde aqui o desde parámetro
	adminRoles := []string{"ADMIN"}
	roles, err := s.RoleRepo.GetByNames(ctx, adminRoles)
	if err != nil {
		return err
	}

	// * verificar quienes usan Roles, para que usen RolesModel
	userCreate.RolesModel = roles
	// * en todos voy a poner los dos
	//auth.UserAgregate.Roles = adminRoles

	err = s.AuthTx.SignUpUser(ctx, userCreate)
	if err != nil {
		return err
	}

	// enviar email?

	// igual al sign up verificar si existe el user y el auth con emil password
	// asignar roles - por último injectas el servicio de sign up por que es igual
	// * usando sin ser instanceado y asignando
	//auth.UserAgregate.CreatedAt = &now

	// **

	return nil
}
