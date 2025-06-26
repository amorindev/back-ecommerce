package postgresql

import (
	"context"
	"errors"

	"com.fernando/pkg/app/auth/model"
)

// * jsonscheemma
// * como mantener actualizado igual al embedding de mongo
// * transacciones si se usa embedign en postgresql
// * mi recomendación pasar todas las tablas de postgresql a mongo, si necesitas un campo adicional
// * la tabla generada de muchos  muchos el tiempo o quien lo creo el id no me veo
// * en la necesidad de modiificar toda mis estrutura,
// * lo vas puliendo si sabes que no va ser necesario embeding la tabla que se genera
// * esto nos ayuda a la lectura, pero en la escritura tenemos algunos pasos adicionales
// * Falta en get buscar mediante relaciones se esat usando el embeding
// * twitter o X mejor el cantidad delikens
// https://www.youtube.com/watch?v=HyIDdvKX3VQ

// !modelar a base de datos

// * Sacrificar disco por salvar CPU, depende del contexto, del tipo de app
// ! reponsabilidad mantener los aatos actualizados en los json schemma, o embedding
// * Asignar role - auth changes - roles y  permisison en el token o en el User?
// * a que entidad afecta


func (r *Repository) CreateEbd(ctx context.Context, auth *model.Auth) error {
	return errors.New("auth postgresql repo - CreateEdb unimplement")
}
