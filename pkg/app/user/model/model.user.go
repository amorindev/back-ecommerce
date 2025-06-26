package model

import (
	"time"

	"com.fernando/pkg/app/auth/model"
	roleModel "com.fernando/pkg/app/role/model"
	"com.fernando/pkg/app/user/constants"
)

// * el señor que publica en facebook guarda la ip ver para bloquear cantidad de request
// * y la session que incluya el plataforma si es mobile o web desde donde se creo
// * User agregate
// TODO: implementar
// * ver si coliciona con User de user servicio
// * deberia dejaro como puntero o sin puntero
// * para validar
// * poner intrface{} tiene una gen responsabiidad a obtener el id parsearlo o no?
// * verificar si al parcear con Hex, es igual a lo que hace json cuando responde, lo hace implicitamente
// * UserAgregate o User
// que se debe tener encuenta si uso []*role.Role puntero
// agregar phone, verified?
// añl hacer todos lo campos punteros por que se insertrán en la base de datos como será el flujo
// gregar phone - retornará el phone por default no todos los phones
// el photo url del bucket si es sign in  y sign up
// si es con provider y tiene photo
// lo traemos cul es el flujo se guada en labase de datos o se obtiene al hacer sign in
// * RolesIDs, parece para mongo
// * add bson esto me ayudaria  como aux ára no hacer varias consultas
// * en phone debe salir el phone por defecto me parece mejor set por defecto
// * el ultimo agregado ver
// ? Roles field, o sacarlo desde el token ver
// RolesModel - no se para que es ver no o agregar otra capa // add amitempty
// RolesModel - // ? es puntero como manejarlo en las capas(buscar), "-" puede ser omitempty?
// como manejar la imagen de sing in consocial medias y nuestra imagen de usar
// al agregar verificar que si no existe una imagen en user ponerlo de sign in provider
// o no se
// ImgUrl como cargar la imagen ver si agregar los dos campos aqui o separarlo
// las imagenes por defecto de daca DDD se guarda en la misma crpeta como default
// o guardarlo aparte en utils mejor en ada carpeta
// desde donde se debe cargar la imagen por defecto desde el env pasandole la url desde donde?
// Phone phoner si si es verificado? ver condiciones
// faltría el usernaem (`json:"usernanme" bson:"username"`)y el adrress
// RolesIDs no se si se va usar en postgresql tambien o sino hacer que se retorne
// RolesIDs add bson esto me ayudaria  como aux ára no hacer varias consultas
// RolesModel seria solo role por que solo usaremos el name y no todo el objeto
// RolesModel ver si va ser puntero
// ! ver que campos estan sobrecargando la entidad
// falta phone verified
// * go-l solo hace *time.Time y a otras entidades me parece que no esta controlando los
// * me parece que todos los valores son requreridos
// * AuthProviderCreate develiver las lista de proveerores y comose cual usao
// * estrategia de firebase retorna todo ver que se ve a retornear per me gusta
// TODO TwoFaMethod dos cosas si va ser tipo enum o sera relacionado con otra tabla el tipo
// TODO si IsTwoFaEnabled esta en false TwoFaMethod debe ser nil
// RolesModel me parece que si o si se saca por que solo se necesita el name
type User struct {
	ID                 interface{}            `json:"id" bson:"_id,omitempty"`
	Email              string                 `json:"email" bson:"email"`
	EmailVerified      bool                   `json:"email_verified" bson:"email_verified" db:"email_verified"`
	RolesModel         []roleModel.Role       `json:"-" bson:"roles"`
	UserName           *string                `json:"username" bson:"username"`
	Name               *string                `json:"name" bson:"name"`
	Roles              []string               `json:"roles" bson:"-"`
	Phone              *string                `json:"phone" bson:"phone"`
	PhoneVerified      bool                   `json:"phone_verified" bson:"phone_verified"`
	ImgUrl             *string                `json:"img_url" bson:"img_url"`
	AuthProviderCreate *model.Auth            `json:"-" bson:"-"`
	CreatedAt          *time.Time             `json:"created_at" bson:"created_at"`
	UpdatedAt          *time.Time             `json:"updated_at" bson:"updated_at"`
	IsTwoFaEnabled     bool                   `json:"is_two_fa_enabled" bson:"is_2fa_enabled"`
	TwoFaMethod        *constants.TwoFaMethod `json:"-" bson:"two_fa_method"`
	//RolesIDs           []interface{}    `json:"-" `
	//AuthProviders      []*model.Auth    `json:"auth_providers" bson:"auth_providers"`
	//Roles         *[]*roleModel.Role `json:"roles"` // add amitempty // ? es puntero como manejarlo en las capas(buscar)
	//Roles         []*roleModel.Role `json:"roles"` // add amitempty // ? es puntero como manejarlo en las capas(buscar)
	// hacer ejemplos puedes insertar nil en  []*roleModel.Role y dará error de ejecucion
	// al recorrer hacer un if
}

func NewUserSignUp(email string, name *string, username *string, password string, phone *string) *User {
	now := time.Now().UTC()
	return &User{
		Email:         email,
		EmailVerified: false,
		Name:          name,
		Phone:         phone,
		AuthProviderCreate: &model.Auth{
			Password:  password,
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		IsTwoFaEnabled: false,
		TwoFaMethod:    nil,
		// revisar los demas campos
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
/*
func (s *AuthService) SignUp(email string, name *string, password string) (*user.User, error) {
    user := userM.NewUser(email, name)
    auth := authM.NewAuthProvider(password)
    
    user.SetAuthProvider(auth) // ← método interno para encapsular lógica
    
    err := s.tx.SignUpWithOtpUser(user, ...)
    return user, err
}
*/

/*
type UserWithAuth struct {
	User *user.User
	Auth *auth.Auth
}cuando hacer esto
*/

/*
💡 4. ¿En los puertos (interfaces), debo usar solo un modelo por dominio?
✅ Sí, los puertos de auth deberían depender solo de auth, no de user.
En otras palabras:
// auth/port/repository.go
type Repository interface {
	CreateAuth(ctx context.Context, auth *Auth) error
	GetByEmail(ctx context.Context, email string) (*Auth, error)
	...
}
*/