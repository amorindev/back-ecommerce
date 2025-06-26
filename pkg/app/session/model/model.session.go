package model

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// cambiar a sesiones
// mirar el video
// add bson ttags
// igual que el hash password deberia sacar el create hash de user para hacerlo global?
// refresh token deberia encriptarlo como password?
// RefreshTokenID, para generar el refreshtoken necesito el id por eso
// en e caso de mongo no quiero crear el ID desde el servicio esto crearía dependencia
// es por eso que tanto pg  como nmongo crean su ID, y agrego un campo,
// uuid string para ambos pg y mongo
// ? de la session solo se devuelve el access refresh token
// * para cambiar contraseña si devolver el device
// hacer el id *string ? para null
// agregar el proveedor auth?
// ! ES UN AGREGADO DE USER
// ! ExpiresAt
// AccessToken no se guarda
// teniamos el caso de el id RefreshTokenID sirve para buscar no suar el id principal
// por que no podiamos crear
// Email no se sabe para que se usa
// UserRolessirve para crear el token
// puede ser para saber los roles de la session UserRoles
// UserRoles        []string    `bson:"-"` // sirve para crear el token
// si hay
// me parece mejor parsarlo como parámetro RemenberMe
// refresh token si se guarda
type Session struct {
	ID               interface{} `json:"-" bson:"_id"`
	UserID           interface{} `json:"-" bson:"user_id" `
	AccessToken      string      `json:"access_token" bson:"-"`
	RefreshTokenID   string      `json:"-" bson:"refresh_token_id"`
	RefreshToken     string      `json:"refresh_token" bson:"-"` // se guarda el token hashed
	RefreshTokenHash *string     `json:"-" bson:"refresh_token"`
	Device           *string     `json:"-" bson:"device"`
	Revoked          bool        `json:"-" bson:"revoked"`
	Email            string      `json:"-" bson:"-"`
	RemenberMe       bool        `json:"-" bson:"-"` // sirve para crear el token
	ExpiresAt        time.Time   `json:"-" bson:"expires_at"`
	ExpiresIn        int64       `json:"expires_in" bson:"expires_in"`
	CreatedAt        *time.Time  `json:"-" bson:"create_at"`
	// debería tener updated ? se deberia actualizar unregistro
	// o es para el revocked puede ser
	// IP ? - si es mobile?, se puede combinar la limit ratind
}

func (t *Session) HashToken() error {
	tokenHash, err := bcrypt.GenerateFromPassword([]byte(t.RefreshToken), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	tH := string(tokenHash)

	t.RefreshTokenHash = &tH

	return nil
}

func (t *Session) TokenMatch(token string) (bool, error) {
	if t.RefreshTokenHash == nil {
		return false, errors.New("password hash is nil")
	}

	err := bcrypt.CompareHashAndPassword([]byte(*t.RefreshTokenHash), []byte(token))
	if err != nil {
		return false, err
	}

	return true, nil
}
