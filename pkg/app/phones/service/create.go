package service

import (
	"context"
	"time"

	"com.fernando/pkg/app/phones/model"
)

func (s *Service) Create(ctx context.Context, phone *model.Phone) error {
	// ! el number debe ser unique dos formas consultando y agregando contrints
	// !
	//fmt.Printf("===================1\n")
	now := time.Now().UTC()

	phone.CreatedAt = &now
	phone.UpdatedAt = &now
	isverified := false
	phone.IsVerified = &isverified
	// de donde hacerlo por que la logica es que el ultimo agregado sea el por defecto

	// * si funciona pero ver demomento no convinamos
	// ! Este debe ser otra accion del usuario para ello tenemos mark by default y nocombinar todo
	phone.IsDefault = false
	/* phone.IsDefault = true

	p, err := s.PhoneRepo.GetDefault(ctx)
	if err != nil {

		if err != errors.ErrPhoneNotFound {
			return err
		}
	}
	fmt.Printf("Phone: %+v\n", p)

	if p != nil {

		return s.PhoneTx.Insert(ctx, p.ID.(string), phone)
	} */

	return s.PhoneRepo.Insert(ctx, phone)
}
