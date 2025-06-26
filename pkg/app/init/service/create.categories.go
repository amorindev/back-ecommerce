package service

import (
	"context"

	"com.fernando/pkg/app/ecomm/category/errors"
	"com.fernando/pkg/app/ecomm/category/model"
)

func (s *Service) CreateCategories() error {
	// * mejor que solo creo los que falta ver - por que agregue uno y ya no me crea ninguno
	// * o eiminar todas las colleciones y empezar de nuevo en postgresql nos dice que hay una
	// * oreign key aqui no veo delete cascade, se puede alterar los datos
	ctgs := []string{"HOMBRE", "MUJER","UNISEX","NIÑOS"}


	var categories []*model.Category
	for _, c := range ctgs {
		cate, err := s.CategoryRepo.GetByName(context.Background(),c)
		if err != nil {
			if err != errors.ErrCategoryNotFound {
				return err
			}
		}
		if cate !=nil {
			return nil
		}
		var ctg model.Category
		ctg.Name = &c
		categories = append(categories, &ctg)
	}
	
	// TODO: que pasa si agregas agregas created llamarías al servicio para no crearlo directamente
	// ? o desde aqui ver 
	// * lo separo por que quiero asegurarme  que se inserte todos
	err := s.CategoryRepo.CreateMany(context.Background(),categories)
	if err != nil {
	  return err
	}
	return nil
}


