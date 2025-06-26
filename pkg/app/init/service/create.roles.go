package service

import (
	"context"
	"strings"

	"com.fernando/pkg/app/role/errors"
	"com.fernando/pkg/app/role/model"
)

func (s *Service) CreateRoles(ctx context.Context, roleNames []string) error {
	// * asegurar se que se creen todos los roles
	// otra op cion seria crear un arreglo pero pero duplcariamos el for en repository o service
	// crearlo de esta manero o desde repo conun for a Exec, bueno primero tengo que buscarlo
	//capitalice

	// ? se debe crear role entidad desde afuera?
	for _, name := range roleNames {
		var role model.Role
		role.Name = strings.ToUpper(name)
		//_, err := s.RoleRepo.FindByName(ctx, role.Name)
		_, err := s.RoleRepo.GetByName(ctx, role.Name)
		if err != errors.ErrRoleNotFound {
			return err
		}
		err = s.RoleRepo.Insert(ctx,&role)
		if err != nil {
			return err
		}
	}

	/* err := s.RoleRepo.CreateMany(ctx, roleNames)
	if err != nil {
		return err
	} */

	return nil
}

