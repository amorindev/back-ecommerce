package port

import (
	"context"

	"com.fernando/pkg/app/permission/model"
)

type PermissinRepo interface {
	Get(ctx context.Context, id string) (*model.Permission, error)
	Insert(ctx context.Context, permission *model.Permission) error
	GetByName(ctx context.Context, name string) (*model.Permission, error)
	// * me parece que los insert many en todos los DDD son para datos masivos excel
	//* y me parece que usa trnsacciones para ver si existe uno e insertar
	// * y datos de inicio separarlo en otra interfaz?
	InsertMany(ctx context.Context, permissions []*model.Permission) error
	// * retorna el slice de permission que existen para informar al usuario
	ExistOne(ctx context.Context, permissions []*model.Permission) ([]*model.Permission, error)
	// * ver las dos formas de momento ambas válidas en la segunta ya tenemos el id dentro de
	// * la estructura
	AssignPermissionsToRole(ctx context.Context, roleID string, permissions []*model.Permission) error
	// desde el forntend solo vienen los dos ids mejor usar este 
	//AssignPermissionsToRole2(ctx context.Context, permissions []*model.Permission) error
}
