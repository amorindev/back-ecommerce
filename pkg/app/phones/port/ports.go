package port

import (
	"context"

	"com.fernando/pkg/app/phones/model"
)

// * Cuando el usuario inicia session se retorna el user con solo el telefono por defecto
// * tambien se podría agregra el  address por defecto
// * desde mi bakend si nohay un imagne por defecto para el user obtenerlo desde minio y asignarlo ver si separar los default en otro file como utils o en el mismo DDD user/img_default.png
type PhoneRepo interface {
	Insert(ctx context.Context, phone *model.Phone) error
	Get(ctx context.Context, id string) (*model.Phone, error)
	GetAll(ctx context.Context, userID string) ([]*model.Phone, error)
	GetDefault(ctx context.Context) (*model.Phone, error)
	ChangeDefault(ctx context.Context, id string, isDefault bool) error
}

type PhoneSrv interface {
	Create(ctx context.Context, phone *model.Phone) error
	GetAll(ctx context.Context, userID string) ([]*model.Phone, error)
	ChangeDefault(ctx context.Context, id string, isDefault bool) error
}

type PhoneTransaction interface {
	// * necesitamoe el phone id del anterior que fue marcado como true
	Insert(ctx context.Context, phoneDefaultID string, phone *model.Phone) error
}
