package service

import (
	"context"
	"time"

	"com.fernando/pkg/app/ecomm/address/model"
)

func (s *Service) Create(ctx context.Context, address *model.Address) error {
	// ? verificar que el usuario exista, y crear existe user bool error
	// ver las validaciones errUsernotfound si elerror is nil
	now := time.Now().UTC()

	address.CreatedAt = &now
	address.UpdatedAt = &now

	// * igual para phones
	// * al crear un address se debe marcar como por defecto desde aqui me parece bien
	// * ver donde hacerlo desde mobile documenta
	address.IsDefault = false

	// ! Este debe ser otra accion del usuario para ello tenemos mark by default y nocombinar todo
	/* address.IsDefault = true

	// * se necesita una transaccion

	add, err:= s.AddressRepo.GetDefault(ctx)
	if err != nil {
	  if err!= errors.ErrAddressNotFound {
		return err
	  }
	}
	if add != nil {
		return s.AddressTx.Insert(ctx, add.ID.(string),address)
	} */

	return s.AddressRepo.Insert(ctx, address)
}
