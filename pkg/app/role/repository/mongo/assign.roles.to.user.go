package mongo

import (
	"context"
	"fmt"
	
	"com.fernando/pkg/app/role/model"
	mongoPkg "go.mongodb.org/mongo-driver/v2/mongo"
	
	"go.mongodb.org/mongo-driver/v2/bson"
)

// llamarlo assign o add
// no es mejor usar []model.UserRole
// Alternativa en MongoDB sin transacción: Usa InsertMany, que falla en bloque si algún documento no puede insertarse.
// ? se debe usar una transaccion en InsertMany?
func (r *Repository) AssignRolesToUser(ctx context.Context, userID string, roles []model.Role) error {
	userObjID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("role mongo repo - AssignRolesToUser err: %w", err)
	}

	//var roleUsers []interface{}
	var roleUsers []bson.M

	for _, role := range roles {
		// ! primero verificar si  existe esa relación  ya si con todos

		roleObjID, err := bson.ObjectIDFromHex(role.ID.(string))
		if err != nil {
			return fmt.Errorf("role mongo repo -AssignRolesToUser err: %w", err)
		}
		count, err := r.UserRoleColl.CountDocuments(ctx, bson.M{"user_id": userObjID, "role_id": roleObjID})
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}
		roleUsers = append(roleUsers, bson.M{
			"user_id": userObjID,
			"role_id": roleObjID,
		})
	}

	session, err := r.Client.StartSession()
	if err != nil {
		return fmt.Errorf("role mongo repo - AssignRolesToUser err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("role mongo repo - AssignRolesToUser err: %w", err)
	}

	defer session.EndSession(ctx)

	err = mongoPkg.WithSession(ctx, session, func(ctx2 context.Context) error {
		// * dentro de aui se usaría el context 2
		// ? debería ir aqui session.StartTransaction
		_, err := r.UserRoleColl.InsertMany(ctx2, roleUsers)
		if err != nil {
			return err
		}

		return session.CommitTransaction(ctx2)
	})
	if err != nil {
		if err := session.AbortTransaction(context.Background()); err != nil {
			return fmt.Errorf("role mongo repo - AssignRolesToUser abort transaction err: %w", err)
		}

		return fmt.Errorf("role mongo repo - AssignRolesToUser err: %w", err)
	}

	return nil
}

