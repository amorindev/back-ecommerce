package service

import (
	"context"

	"com.fernando/pkg/app/ecomm/products/model"
)

func (s *Service) Get(ctx context.Context) ([]*model.Product, error) {
	products, err := s.ProductRepo.Get(ctx)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		product.CategoryID = nil

		presignedUrl, err := s.FileStorageSrv.GetImageUrl(ctx, product.FileName)
		if err != nil {
			return nil, err
		}
		product.ImgUrl = presignedUrl.String()

		for _, productItem := range product.ProductItems {
			presignedProductItemURL, err := s.FileStorageSrv.GetImageUrl(ctx, productItem.FileName)
			if err != nil {
				return nil, err
			}
			productItem.ImgUrl = presignedProductItemURL.String()
		}

	}

	return products, nil
}
