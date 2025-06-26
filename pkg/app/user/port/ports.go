package port

import (
	"context"

	phoneM "com.fernando/pkg/app/phones/model"
	"com.fernando/pkg/app/user/constants"
	userM "com.fernando/pkg/app/user/model"
)

// ! como manejar los embedding agregar otra función  y en la otras capas como validarlo
// * si deseas ingresar a user que es interno a auth no sea null puntero
// ? devería guardar el email en user tambien - habria mejora en rendimiento?

// * como los crud de auth get delete update son para cada repo por
// * ejemplo al delete account debes eliminar auth mediante su repo user mediante su repo en una transaccíón

// * ver como funciona supabase y firebase para cundo el usario quiere agregar su propia informacín del profile
// * separarlo completamente o el programdor use User{} y agregue en campo Profile o Profile agregate

// * como manejar el current user cuando el solo auth o tambien user y profile, cuando uso AssignRoles()

// * como se comportan con role user auth profile y las transacciones

// /DOS formas injectando el UserRepo o compartiendo en user shared
type UserSrv interface {
	// el update si puede tener  por que no esta tan ligado

	GetAll(ctx context.Context) ([]*userM.User, error)
	GetOne(ctx context.Context, id string) (userM.User, error)
	GetByEmail(ctx context.Context, email string) (*userM.User, error)
	GetUser(ctx context.Context, userID string) (*userM.User, error)
	//Create(ctx context.Context, user *User) error // esto es igual que register
	//este Auth update diferencia - esta relacionado con modificar los datos del usuario comodireccion u otros
	//modifica nombre apillidos roles, se maneja incuyendo roles y permisos
	Update(ctx context.Context, id string, user userM.User) error
}

// ! asegurarse de usar el context de a trnasaccion en cada repo mongo o postgresql
type UserRepo interface {
	Get(ctx context.Context, id string) (*userM.User, error)
	// como manejarlo con ebd o crear otra función, los anidados roles
	GetByEmail(ctx context.Context, email string) (*userM.User, error) //no es necesario anidar elusuario?
	Insert(ctx context.Context, user *userM.User) error
	CreateEbd(ctx context.Context, user *userM.User) error
	ConfirmEmail(email string) error
	EnableTwoFaSms(ctx context.Context, userID string, twoFaMethod constants.TwoFaMethod) error
	// * de momento en deleteAccount service eliminaremos el user
	Delete(ctx context.Context, id string) error // se relaciona con Delete account de AuthService, transaccion

	// como se combina delete con auth providers delete y delete permanentemente

	// GetAll y GetOne obtienen información basica del usario sin incluir password como en microblog
	// ver si se va usar puntero en ambos casos
	//GetAll(ctx context.Context) ([]model.User, error)
	//GetOne(ctx context.Context, id string) (model.User, error)

	//Update(ctx context.Context, id string, user User) error
	//Delete(ctx context.Context, id string) error

	//AssignRole(ctx context.Context, userID string, roleID string) error

	// !ver si se va a separar
	InsertTwoFaSms(ctx context.Context, twoFaSms *userM.UserTwoFaSms) error
	// el phone asignado a twoFaSms usando join no puede ser get a phone colleciton
	// por que necesitamos el id de la relacion
	// * estoy retornando Phone, que? o solo retorno el número ver
	// obtener el phone asiciado al twoFaSms
	// ? dende deveria estar esta funcion ahora phone se relacion con UserTwoFaSms
	// ver comoquedara las  y como asegurarse que se unico ver tambien en postgresql
	GetPhonetwoFaSms(ctx context.Context, userID string) (*phoneM.Phone, error)
}

//Si los roles están estrechamente relacionados con la autenticación, podrías considerar mantener
// FindWithRoles en AuthRepository. Si los roles son una entidad independiente y se usan en múltiples
// contextos, entonces inyectar RoleRepository o RoleService en AuthService es una excelente opción.
// USAR CASCADE? o transacciones, si me paso a mongo

//TODO: Auth User
// TODO: Campos nulos: Si un campo puede ser nulo, es recomendable usar punteros
// TODO: (por ejemplo, *string, *time.Time). Esto permite distinguir entre un campo que
// TODO: no se ha establecido y un campo que se ha establecido explícitamente como nulo.
// TODO: como validaros para que no cause errores

// * Roles: Si los roles son una parte importante de la entidad User,
// * es recomendable manejarlos dentro de UserRepository. Si los roles
// * tienen una lógica compleja, podrías considerar crear un RoleRepository separado.

// TODO: ConfirmPassword: Este campo no debería ser almacenado en la base de datos,
// TODO: ya que es solo para validación en el momento del registro o cambio de contraseña.
// TODO: Por lo tanto, no es necesario que sea un puntero. (verificar hashpassword )
// TODO: parece que de orlando estaba bien

// * Provider: Si el proveedor de autenticación (por ejemplo, Google, Facebook)
// * es opcional, es correcto que sea un puntero.

// si debeovemos []*Entity es por si queremos cambiar los datos  o hacer calculos ?
