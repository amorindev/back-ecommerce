package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/role/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// ? []*model.Role si se rá modificado en otra capa, o si no quiero crear con valores por defecto en
// ? el caso de una strct{}
// ! revisar CreateMany de categorias
func (r *Repository) FindByNamesRole(ctx context.Context, names []string) ([]model.Role, error) {

	var roles []model.Role

	for _, roleName := range names {
		// ! Donde esta el id
		//id := bson.NewObjectID()
		// !------------------
		var role model.Role
		err := r.Collection.FindOne(context.TODO(), bson.D{{Key: "name", Value: roleName}}).Decode(role)
		if err != nil {
			// ! Refactorizar - que pasa si el  erro de no documents
			return nil, fmt.Errorf("role mongo repo - FindByNamesRole error: %w", err)
		}
		roles = append(roles, role)
	}
	// parsear ids

	return roles, nil
}
