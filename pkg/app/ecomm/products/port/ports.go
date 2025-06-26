package port

import (
	"context"

	"com.fernando/pkg/app/ecomm/products/model"
)

type ProductRepo interface {
	Insert(ctx context.Context, product *model.Product) error
	// * por defecto es con paginación,
	GetAll(ctx context.Context, limit int, page int) ([]*model.Product, error)
	GetByName(ctx context.Context, name string) (*model.Product, error)
	//Get(ctx context.Context, id string) (*model.Product, error) agregar esto
	// * seria mejor obtener mediante el sku, sku como id? o el dni me parece que no
	// * mejor ambos por que igual mongo crea el id
	//GetBySku(ctx context.Context, sku string) (*model.Product, error)
	Get(ctx context.Context) ([]*model.Product, error)
}

// TODO comparar con los de product
// Create(ctx context.Context, product *model.Product) error
// CreateWithVariants(ctx context.Context, product *model.Product) error
type ProductSrv interface {
	Create(ctx context.Context, product *model.Product) error
	Get(ctx context.Context) ([]*model.Product, error)
	/* deepseek
	response := &GetAllResponse{
		Page:     page, no veo que se usa a no ser que sea la pagina actual y se pude savar del
		mismo handler  no creo que en el proceso se cambie
		Limit:    limit,
	}
	*/
	/// count y pages
	// !como sacarlos demomento uso de la  request
	GetAll(ctx context.Context, limit int, page int) ([]*model.Product, int, int, error)
}

type ProductTransaction interface {
	Create(ctx context.Context, product *model.Product) error
}
