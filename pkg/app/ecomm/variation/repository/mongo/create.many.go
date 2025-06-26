package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/variation/errors"
	"com.fernando/pkg/app/ecomm/variation/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// es necesario pasar losid a string? ver gosenior ultimo videos
func (r *Repository) CreateMany(ctx context.Context, variations []*model.Variation) error {
	// una opcion sería encapsuar todo esto dentro de la trnasaccion pero loa ¿haré simple
	// ! esto no se hacve aqui se hace desde algun servicio
	for _, v := range variations {
		// *verificar si existe la categoría - o es mejor hacerlo desde repository
		// * hasta ahora es mejor por aqui por que usare transacciones
		// ? debería verificar si es nulo donde hacerlo, *c, ver las validaciones de go senior
		// estamos llamando a una funcion del mismo repo, un repo llamando otro repo
		_, err := r.GetByName(ctx, *v.Name)
		if err != nil {
			if err != errors.ErrVariationNotFound {
				return err
			}
		}
		// * creamos el id
		id := bson.NewObjectID()
		v.ID = id
	}

	session, err := r.Client.StartSession()
	if err != nil {
		return fmt.Errorf("variation mongo repo - CreateMany err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("variation mongo repo - CreateMany err: %w", err)
	}

	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(ctx2 context.Context) error {
		// * dentro de aui se usaría el context 2
		_, err := r.Collection.InsertMany(ctx2, variations)
		if err != nil {
			return err
		}

		return session.CommitTransaction(ctx2)
	})
	if err != nil {
		if err := session.AbortTransaction(context.Background()); err != nil {
			return fmt.Errorf("variation mongo repo - CreateMany abort transaction err: %w", err)
		}

		return fmt.Errorf("variarion mongo repo - CreateMany err: %w", err)
	}

	// Dos opciones *[]* no podría ser por que no afecta a la ongitud, en la capa de
	// servicio se debería ver reflejado sin problemas usando []*
	// si no retornar el slice
	return nil

}


