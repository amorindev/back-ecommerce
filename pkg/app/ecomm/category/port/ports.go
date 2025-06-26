package port

import (
	"context"

	"com.fernando/pkg/app/ecomm/category/model"
)

type CategoryRepo interface {
	Insert(ctx context.Context, category *model.Category) error
	CreateMany(ctx context.Context, categories []*model.Category) error
	GetByName(ctx context.Context, name string) (*model.Category, error)
	GetAll(ctx context.Context) ([]*model.Category, error)
}

type CategoryService interface {
	GetAll(ctx context.Context) ([]*model.Category, error) 
}

/* type CategoryDataLoader interface{

} */
