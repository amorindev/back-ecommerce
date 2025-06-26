package mongo

import (
	"context"
	"errors"

	"com.fernando/pkg/app/auth/model"
)

// ! que pasa si el id es interface verificar lo que genera el Hex funcion y el  formado de encoder json
//
//	Get(id string) (player *domain.Player, err error), go-l
//
// si el auth se va modificar por ejemplo desde servicio comviene un puntero, hacer pruebas sin y con
func (r Repository) Update(ctx context.Context, id string, auth model.Auth) error {
	return errors.New("auth mongo repo - update unimplement")
}
