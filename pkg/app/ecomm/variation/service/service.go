package service

import "com.fernando/pkg/app/ecomm/variation/port"

var _ port.VariationService = &Service{}

// * adt adapter
type Service struct {
	VariationRepo port.VariationRepo
}

func NewService(variationRepo port.VariationRepo) *Service {
	return &Service{
		VariationRepo: variationRepo,
	}
}
