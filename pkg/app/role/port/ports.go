package port

import (
	"context"

	"com.fernando/pkg/app/role/model"
)

type RoleService interface {
}

// Usa roleIds cuando solo necesitas asociar roles existentes a un usuario.
// Usa el objeto Role cuando necesitas crear nuevos roles o realizar operaciones
//
//	que dependen de los datos completos del rol.
//
// Combina ambos si tu lógica requiere tanto la creación de nuevos roles
//
//	como la asignación de roles existentes.
//
// * me parece necesario retornar un arrgloe de nombres de roles
// * para sign in y para el token, en sign in no es necesario retornar el id del role
// * omitempty por que cuando hago get roles si necesito el id,
// * verificar el tema de modular
// * cuando devover la lista
// ! En que casos se debería parsear los IDs estan en varias funciones ver
type RoleRepo interface {
	// * de momento la tabla intermedia en ambos
	// solo id del user y roles
	// Create o asign roles
	AssignRolesToUser(ctx context.Context, userID string, roles []model.Role) error

	// Create many o createSlice
	Insert(ctx context.Context, role *model.Role) error
	//createbynames
	CreateMany(ctx context.Context, names []string) error // []*modelRole vs names, como seria en signup
	GetByName(ctx context.Context, name string) (*model.Role, error)
	GetByNames(ctx context.Context, names []string) ([]model.Role, error)
	//cambiar a get o eliminar
	// esto deberia retornar la entidad usuario lista de entidades roles, se puede crear
	// otro para devolver la entidad esto lo uso para retornarlo directamente al token sin
	// hhacer for, anocer que debuelvas los roles con el id en el token ver
	// ver si sirve hasta ahora para e token que quier unareglo de string
	// que pasa si tenfdría que hacer un for para convertir o hacer otra consulta
	// o pasarle e slice de role entity crearía dependencia que hay de modular?
	GetByUserID(ctx context.Context, userID string) ([]string, error) // vs
	GetByUserIDRole(ctx context.Context, userID string) ([]model.Role, error)

	// * crearlo aparte
	// * Dos opciones si ponemos en RoleRepository depende de Role, role si pasamos a mongo podemos reutilizar
	// * Si esta aqui cumple con princio de single responsability

	// ? me parece al hacer up date como funciona, como funciona junto al update al checbox
	// o la lista de roles ver junto con la interfaz de usuario
	// el administrador desde puede asignar roles a un usurio update roles
	// desde el midleawre injectandolo o desde el service ?
	RemoveRolesToUser(ctx context.Context, userID string, roleID string) error
}
