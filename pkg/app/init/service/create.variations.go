package service

import (
	"context"

	ErrVariation "com.fernando/pkg/app/ecomm/variation/errors"
	"com.fernando/pkg/app/ecomm/variation/model"
)

func (s *Service) CreateVariations() error {
	variationsNames := []string{"COLOR", "SIZE"}

	var variations []*model.Variation
	for _, c := range variationsNames {
		vrt, err := s.VariationRepo.GetByName(context.Background(), c)
		if err != nil {
			if err != ErrVariation.ErrVariationNotFound {
				return err
			}

		}

		if vrt != nil {
			return nil
		}
		var variation model.Variation
		variation.Name = &c
		variations = append(variations, &variation)
	}

	err := s.VariationRepo.CreateMany(context.Background(), variations)
	if err != nil {
		return err
	}

	return nil
}
