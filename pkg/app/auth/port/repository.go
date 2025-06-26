package port

import (
	"context"

	"com.fernando/pkg/app/auth/model"
)

// ! Get Get by email obtendra el id del usario el cual es interface hacer el parceo a string antes de
// ! Eliminar AuthUseModel de los repositories
// ! E único modelo importado debe ser Auth
// acabar con a función
// ? que pasa si esta nulo el embeding
type AuthRepo interface {
	// * Recuer losge t de auth el agreagate es nil
	Get(ctx context.Context, id string) (*model.Auth, error)
	// * Obtiene el auth mediante el ID filtrado por el provider
	// get by userID and provider
	GetByIDProvider(ctx context.Context, userID string, provider string) (*model.Auth, error)

	// @deprecated
	//GetByEmailEbd(ctx context.Context, email string) (*model.Auth, error)

	// ------
	Insert(ctx context.Context, auth *model.Auth) error // es necesario anidar el usuario
	CreateEbd(ctx context.Context, auth *model.Auth) error
	//CreateWithUser(ctx context.Context, auth *model.Auth) error
	//CreateWithUserEbd(ctx context.Context, auth *model.Auth) error
	// * ver como sera delete account de momento eliinaremos solo el usuario en user DDD
	//DeleteAccount(ctx context.Context) error

	// ------
	//update esta relacionado con cambiar el email verified contraseña ,tokens, y el updata de User con
	//datos del usuario
	Update(ctx context.Context, id string, user model.Auth) error

	// ver los campos nulos desde el service como go-l?
	//UpdatePass(ctx context.Context)
	//Delete(ctx context.Context, id string) error
}
