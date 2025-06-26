package service

import (
	"context"

	"com.fernando/pkg/app/ecomm/products/model"
)

func (s *Service) GetAll(ctx context.Context, limit int, page int) ([]*model.Product, int, int, error) {
	// Validación básica
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// TODO cambiar a getall el nombre de la funcion
	products, err := s.ProductRepo.GetAll(ctx, limit, page)
	if err != nil {
		return nil, 0, 0, err
	}

	// En una implementación real, deberías obtener el total de documentos
	// total, err := s.repo.Count(ctx), ademas se posria mejorar al igual que midu
	// agregaba campo nro_seguidores al momento de seguir a un usuario en su base de datos
	// lo mismo tener un campo o tabla donde no sea necesario ir contando cada ves si no al crear
	// ver las reponsabilidades por que si delete el producto lo mismo  para la red socil
	// tambien se debe reflejar CRUD reflejado

	// * Asignar imágenes
	for _, product := range products {
		url, err := s.FileStorageSrv.GetImageUrl(ctx, product.FileName)
		if err != nil {
			return nil, 0, 0, err
		}

		// ! falta segun variables de entorno
		product.ImgUrl = url.String()

		for _, product := range product.ProductItems {
			productURL, err := s.FileStorageSrv.GetImageUrl(ctx, product.FileName)
			if err != nil {
				return nil, 0, 0, err
			}
			// !falta cambiar dependiendo de la variable de entorno
			product.ImgUrl = productURL.String()
		}
	}

	return products, 820, 35, nil
}
