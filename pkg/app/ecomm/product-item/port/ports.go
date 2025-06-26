package port

import (
	"context"

	"com.fernando/pkg/app/ecomm/product-item/model"
)

type ProductItemRepo interface {
	Insert(ctx context.Context, productItem *model.ProductItem) error
	CreateMany(ctx context.Context, productItems []*model.ProductItem) error
}

type ProductSrv interface {
	
}
