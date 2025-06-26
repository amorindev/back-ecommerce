package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/category/errors"
	"com.fernando/pkg/app/ecomm/category/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// * tendría que retornarlo ? []* por qu solo afecta al objetoy no a la lista
// * de momento solo necesito que se inserte correctamente
func (r *Repository) CreateMany(ctx context.Context, categories []*model.Category) error{
	// una opcion sería encapsuar todo esto dentro de la trnasaccion pero loa ¿haré simple
	for _, c := range categories {
		// *verificar si existe la categoría - o es mejor hacerlo desde repository
		// * hasta ahora es mejor por aqui por que usare transacciones
		// ? debería verificar si es nulo donde hacerlo, *c, ver las validaciones de go senior
		// estamos llamando a una funcion del mismo repo
		_, err := r.GetByName(ctx,*c.Name)
		if err != nil {
			if err != errors.ErrCategoryNotFound {
				return err
			}
		}
		// * creamos el id
		id := bson.NewObjectID()
		c.ID = id

	}
	
	session, err := r.Client.StartSession()
	if err != nil {
		return fmt.Errorf("category mongo repo - CreateMany err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("category mongo repo - CreateMany err: %w", err)
	}

	defer session.EndSession(ctx)


	err = mongo.WithSession(ctx, session, func(ctx2 context.Context) error {
		// * dentro de aui se usaría el context 2
		_, err := r.Collection.InsertMany(ctx2, categories)
		if err != nil {
			return err
		}

		return session.CommitTransaction(ctx2)
	})
	if err != nil {
		if err := session.AbortTransaction(context.Background()); err != nil {
			return fmt.Errorf("category mongo repo - CreateMany abort transaction err: %w", err)
		}

		return fmt.Errorf("category mongo repo - CreateMany err: %w", err)
	}

	// Dos opciones *[]* no podría ser por que no afecta a la ongitud, en la capa de
	// servicio se debería ver reflejado sin problemas usando []*
	// si no retornar el slice
	return nil
}

// * Ver cual elegir que me pasen el modelo sería  lo mas optimo
/* func (r *Repository) FindByNamesRole(ctx context.Context, names []string) ([]model.Role, error) {
} */

// * mira este otro lo pasa al bson.M{}
/* func (r *Repository) AssignRolesToUser(ctx context.Context, userID string, roles []model.Role) error {
	

} */

