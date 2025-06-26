package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/permission/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) AssignPermissionsToRole(ctx context.Context, roleID string, permissions []*model.Permission) error {
	roleObjID, err := bson.ObjectIDFromHex(roleID)
	if err != nil {
		return fmt.Errorf("permission mongo repo - AssignRoles err 1: %w", err)
	}

	var roleUsers []bson.M

	for _, permission := range permissions {
		// ! primero verificar si  existe esa relación  ya si con todos
		permissionObjID, err := bson.ObjectIDFromHex(permission.ID.(string))
		if err != nil {
			return fmt.Errorf("permission mongo repo -AssignRolesToUser err: %w", err)
		}
		count, err := r.RolePermissionColl.CountDocuments(ctx, bson.M{"role_id": roleObjID, "permission_id": permissionObjID})
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}
		roleUsers = append(roleUsers, bson.M{
			"role_id": roleObjID,
			"permission_id": permissionObjID,
		})
	}

	session, err := r.Client.StartSession()
	if err != nil {
		return fmt.Errorf("permission mongo repo - AssignRolesToUser err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("permission mongo repo - AssignRolesToUser err: %w", err)
	}

	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(ctx2 context.Context) error {
		// * dentro de aui se usaría el context 2
		// ? debería ir aqui session.StartTransaction
		_, err := r.RolePermissionColl.InsertMany(ctx2, roleUsers)
		if err != nil {
			return err
		}

		return session.CommitTransaction(ctx2)
	})
	if err != nil {
		if err := session.AbortTransaction(context.Background()); err != nil {
			return fmt.Errorf("permission mongo repo - AssignRolesToUser abort transaction err: %w", err)
		}

		return fmt.Errorf("permission mongo repo - AssignRolesToUser err: %w", err)
	}

	return nil
}


