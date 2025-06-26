package service

import "com.fernando/pkg/app/ecomm/category/port"

var _ port.CategoryService = &Service{}

type Service struct{
	CategoryRepo port.CategoryRepo
}

func NewService(categoryRepo port.CategoryRepo) *Service{
	return &Service{
		CategoryRepo: categoryRepo,
	}
}