package port

import (
	"context"

	addressM "com.fernando/pkg/app/ecomm/address/model"
	"com.fernando/pkg/app/ecomm/stores/model"
)

// ! Recuerda si es ge tby name o title o email deben ser unicos
// ! en la base dedatos o asegurarse de que no exista al momento de insertar
type StoreRepo interface {
	// si se desea cargar desde
	Insert(ctx context.Context, store *model.Store) error
	// que se llame getAll despues se ve si se tiene que agregar paginacion
	GetAll(ctx context.Context) ([]*model.Store, error)
	// name debe ser unique
	GetByName(ctx context.Context, name string) (*model.Store, error)
}

type StoreSrv interface {
	Create(ctx context.Context, store *model.Store, address *addressM.Address) error
	//InsertMany(ctx context.Context, )
	GetByName(ctx context.Context, name string) (*model.Store, error)
	GetAll(ctx context.Context) ([]*model.Store, error)
}

type StoreTx interface {
	// despues vemos si va address dentro de store
	Create(ctx context.Context, store *model.Store, address *addressM.Address) error
}
