package service

import (
	"context"

	"com.fernando/pkg/app/ecomm/variation-option/model"
)

func (s *Service) CreateVariationOptions() error {
	// ? crear unafuncio util reutilizable?
	//colorVariationOptions := []string{"Rojo", "Blanco", "Negro", "Azul", "Amarillo"}
	colorVariationOptions := []string{"RED", "BLACK","WHITE", "LEAD"}
	//sizeVariationOptions := []string{"5", "8", "10", "36", "38", "40", "XS", "S", "M", "L", "XL"}
	//sizeVariationOptions := []string{"S", "M", "L", "XL"}
	sizeVariationOptions := []string{"36", "37", "38", "39","40","41","42","43","44"}

	var variationOptions []*model.VariationOption

	colorVariation, err := s.VariationRepo.GetByName(context.Background(), "COLOR")
	if err != nil {
		return err
	}
	sizeVariation, err := s.VariationRepo.GetByName(context.Background(), "SIZE")
	if err != nil {
		return err
	}

	for _, vOtp := range colorVariationOptions {
		var variationOpt model.VariationOption
		variationOpt.Value = &vOtp
		variationOpt.VariationID = colorVariation.ID.(string)
		variationOptions = append(variationOptions, &variationOpt)
	}

	for _, sOtp := range sizeVariationOptions {
		var variationOpt model.VariationOption
		variationOpt.Value = &sOtp
		variationOpt.VariationID = sizeVariation.ID.(string)
		variationOptions = append(variationOptions, &variationOpt)
	}

	varOptionSlice, err := s.VarOptionRepo.ExistOne(context.Background(), variationOptions)
	if err != nil {
		return err
	}
	if len(varOptionSlice) > 0 {
		// *deberia retornar o imprimer los que existen
		return nil
	}
	err = s.VarOptionRepo.CreateMany(context.Background(), variationOptions)
	if err != nil {
		return err
	}

	return nil
}
