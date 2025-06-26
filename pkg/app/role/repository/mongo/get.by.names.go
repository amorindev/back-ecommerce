package mongo

import (
	"context"
	"errors"

	"com.fernando/pkg/app/role/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) GetByNames(ctx context.Context, names []string) ([]model.Role, error) {

	filter := bson.M{"name": bson.M{"$in": names}}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	// * como se que debo cerrar como request body la base de datos
	defer cursor.Close(ctx)

	var roles []model.Role
	if err = cursor.All(ctx, &roles); err != nil {
		return nil, err
	}

	for i, role := range roles {

		objID, ok := role.ID.(bson.ObjectID)
		if !ok {
			return nil, errors.New("role mongo repo - GetByNames, failed to conver ID to ObjectID")
		}
		role.ID = objID.Hex()
		roles[i] = role
	}

	if len(names) != len(roles) {
		return nil, errors.New("role mongo repo - GetByNames, names and roles diferente longitud")
	}

	return roles, nil
}
