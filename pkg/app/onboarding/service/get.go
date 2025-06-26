package service

import (
	"context"

	"com.fernando/pkg/app/onboarding/model"
)

func (s *Service) Get(ctx context.Context) ([]*model.Onboarding, error) {
	onboardings, err := s.OnboardingRepo.Get(ctx)
	if err != nil {
		return nil, err
	}

	for _, onboarding := range onboardings {
		presignedUrl, err := s.FileStorageSrv.GetImageUrl(ctx, onboarding.FileName)
		if err != nil {
			return nil, err
		}
		// ya que quitamos el parseo si es  env o prod me parece que solo debe retoranar el string y listo
		onboarding.ImgUrl = presignedUrl.String()
		// * File name tiene tag json para el initdata usaremos omitempty para no retornarlo
		// * ver para crear estructuras aparte y no complique otros partes de código
		// * filename si o si no se retorna
		onboarding.FileName = ""
		// * el expires at se envia para lo  que es admin pero este es un backend de mobile
		// * ver eso en Golang y Flutter, lo voy adejar se puede crear bilioteca reutilizable
		// * y manejarlo no se develvera expires at en mobile docs
	}

	return onboardings, nil
}
