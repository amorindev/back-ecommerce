package port

import (
	"context"

	"com.fernando/pkg/app/ecomm/address/model"
)

type AddressRepo interface {
	// * se necesita marcar por defecto y no solo insertar
	Insert(ctx context.Context, address *model.Address) error
	Get(ctx context.Context, id string) (*model.Address,error)
	// se que las direcciones son pocas no se si debe tener paginacion
	// get es para traer solo 1 - user id desde el token
	GetAll(ctx context.Context, userID string) ([]*model.Address, error)
	GetDefault(ctx context.Context) (*model.Address, error)
	// habra get by id? como sera el get by name o title  o no habrá, retornar el objeto actualizado
	ChangeDefault(ctx context.Context, id string, isDefault bool) error
}

type AddressSrv interface {
	Create(ctx context.Context, address *model.Address) error
	GetAll(ctx context.Context, userID string) ([]*model.Address, error)
	ChangeDefault(ctx context.Context, id string, isDefault bool) error
}

type AddressTx interface {
	// * necesitamoe el address id del anterior que fue marcado como true
	Insert(ctx context.Context, addressDefaultID string, address *model.Address) error
}
