package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// *En la capa de servicio deberiamos buscar el user o seria un paso innecesario
// * si no existe user not found y lo llevamos al login eliminando tokens
// * si delete account es igual a 0 mismo proceso solo es informativo y
// * ir a la pantalla de login
// deberia estar en user DDD
func (r *Repository) Delete(ctx context.Context, id string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("auth mongo repo - DeleteAccount err: %w", err)
	}

	filter := bson.M{"_id": objID}

	result, err := r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("user-not-found")
	}

	return nil
}
