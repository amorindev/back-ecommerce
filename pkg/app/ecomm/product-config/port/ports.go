package port

import (
	"context"

	"com.fernando/pkg/app/ecomm/product-item/model"
)

type ProductConfigRepo interface {
	InsertMany(ctx context.Context, products []*model.ProductItem) error
}
