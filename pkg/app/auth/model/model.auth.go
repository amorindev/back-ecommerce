package model

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// * "-" Cuando nunca será retornado (Sign in or Sign up)
// * Confirm password en la función validate del handler o al crear el user Sign up pasarlo sin agregarlo al struct
// * "omitempty", desde el servicio debo vaciarlo a nulo o asegurarme que sea nulo el campo que no
// * quiero mostrar (sign in Sign up)
// * "interface{}" por que puede ser el objeto de mongo uuid o int, asegurarme cambiar el tipo en la
// * capa de repository
// * "*" por que se va agregar a la base de datos - como validarlo ?
// * "sin *" como valdarlo, cuando no usar * como el sever grpc
// ? Passsword y hash password junots?
// deberia ser puntero como se asigna en la capa de servicio?, en email verified u provider
// debveria cambiar la entidad para usar apple o google auth o no?
// revisar los punteros y los -
// ver por que time es puntero por que esnulo?, o por que si no lo ponemos no se modificará en otras capas?
// como validar punteros
// si es mongo bson"-" debe ser puntero? ver por que time es puntero
// depende de tu logica puedes jugar con las etiquetas, al principio puedes tener un mounstruo,
// puedes empesar creando un monstruo o varias entidades y luego ir afinando
// o muchas estrucuras
// graficos como el que pubica en facebook y el tuyo para auth
// mantener en user? -al final sign in es create user o acocunt
// y como no uso db tag en postgresql, me da esa versatilidad de no usar db y hecerlo manualmente ayuda
// ! el id del user bson"-" dependiendo
// refreshtoken debe ser hasseado como el password? y si necesita un campo adicionañ?
// bson"-" solo para guardar?
// bson:",omitempty" → No almacena el campo si está vacío.
// bson:",inline" → Inserta los campos de una estructura embebida directamente.
// ! el email verified esta en user en la base de datos verificar email y email verified
// el auth user es para todo o que tiene que ver con authenticacion
// para que el desarrolador debe crear (actualizar?), debe crear users collections
// falta create_at updated at flujo con base de datos
// falta provider toke o id token, es mejor retorlo por que el público
// tb_auth_providers?
// si es oauth google passwor dis nil
// ! al hacer Auth to proto funcion tambien que se conviertan los gregados user y denotr de user profile
// ! ver que el OtpID no se muestre en el sign in

// ! --------------- quitar access token y refresh son de la session ademas auth debe estar dentro de user
// * falta modelar el accessotoken y refreshtoken en sign in resp
// chatgpt Datos Firebase Supabase
// Provider  esta chevere para saber si es biometric y consultar a otra tabla (ver me parece que no)
// falta asociar cuenta 
type Auth struct {
	ID           interface{}     `json:"-" bson:"_id,omitempty"`           // omitempty, "-"
	UserID       interface{}     `json:"-" bson:"user_id"`                 // !flujo para el userID
	Provider     *string         `json:"provider" bson:"provider"`         // es visible?	
	Password     string          `json:"-" bson:"-"`                       // * es json - por que tenemos otra entidad para sign in y sign up que requieren password además password y passwordHash no se retornan
	PasswordHash *string         `json:"-" bson:"password"`                // *  aux Service
	CreatedAt    *time.Time      `json:"-" bson:"created_at"`
	UpdatedAt    *time.Time      `json:"-" bson:"updated_at"`
	//falta provider id
	// delete cascade si se elimina el user se elimina las cuentas
	// provider unique
}

// ! que pasa si es nil auth dará error en compilcion
// * todos los campos que se insertarán en la base de datos serán puntero, entonces
// * verificar si es nil en PasswordMatch()
// acregar en otra carpeta encript como claim por que sera usado por refresh token
// si es con
func (a *Auth) HashPassword() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	pH := string(passwordHash)

	a.PasswordHash = &pH

	return nil
}

// ver si necesito puntero por que no estoy actualizando
// ! verificar null en el flujo
func (a Auth) PasswordMatch(password string) (bool, error) {
	if a.PasswordHash == nil {
		return false, errors.New("password hash is nil")
	}
	err := bcrypt.CompareHashAndPassword([]byte(*a.PasswordHash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}


/* DDD example
type Car struct {
	ID        string `bson:"_id,omitempty"`
	Name      string `bson:"name"`
	CarDetail CarDetail
}

type CarDetail struct {
	Color    string    `bson:"color"`
	SaleYear time.Time `bson:"sale_year"`
}
*/

/*

func (a *Auth) AuthToProto() (*authpb.Auth, error) {

	id, ok := a.ID.(string)
	if !ok {
		return nil, errors.New("failed to convert Auth to Proto ok")
	}
	return &authpb.Auth{
		Id:       id,
		Email:    a.Email,
		Password: a.Password,
		User: &authpb.User{
			Username: a.User.Name,
			Name:     a.User.Name,
		},
	}, nil
}

func ProtoToUser(a *authpb.Auth) *Auth {
	return &Auth{
		ID:       a.Id,
		Email:    a.Email,
		Password: a.Password,
		User: &User{
			Username: a.User.Username,
			Name:     a.User.Name,
		},
	}
}



func (u *User) ModelToProto() (*authpb.User, error) {
	id, ok := u.ID.(int)
	if !ok {
		return nil, errors.New("ModelToProto - no se pudo convertir a integer")
	}
	newID := int32(id)

	return &authpb.User{
		Id: newID,
	}, nil
}

*/


