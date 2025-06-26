package mongo

import (
	"context"
	"errors"

	"com.fernando/pkg/app/ecomm/product-item/model"
)

func (r *Repository) Insert(ctx context.Context, product *model.ProductItem) error {
	return errors.New("product mongo repo - create unimplement")
}
