package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/role/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// * Cuando el usario administrador asigna usuarios a un role
// * recibir el roleID y retornar los usurios
// * usar transacciones
// debe ir en RoleRepo 90% de acuerdo
// model de make debe tener puntero? revisar la entidad role
// que pasa si se assigna mas de la len() da errors en tiempo de ejecución?
func (r *Repository) AssignRoles(ctx context.Context, userID string, roleIDs []string) error {
	//roleUsers := make([]model.UserRole, len(roleIDs))
	userObjID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("user mongo repo - AssignRoles err 1: %w", err)
	}

	// se puede insert All
	for _, roleID := range roleIDs {
		// ! primero verificar si  existe esa relación  ya si con todos
		urID, err := bson.ObjectIDFromHex(roleID)
		if err != nil {
			return fmt.Errorf("user mongo repo - AssignRoles err 2: %w", err)
		}

		// ! teníamos el problema donde al obtener roles del sign in o como esta función
		// !donde teneoms agreglos de ids entonces crear un auxID
		// id := bson.NewObjectId(); idStr := id.Hex()
		userRole := model.UserRole{
			//! ID: bson.NewObjectID(), por que mongo lo genera automaticamente
			UserID: userObjID,
			RoleID: urID,
		}
		// userRole.ID = idStr

		_, err = r.UserRoleCollection.InsertOne(ctx, userRole)
		if err != nil {
			return fmt.Errorf("user mongo repo - AssignRoles err 3: %w", err)
		}
		//roleUsers = append(roleUsers, userRole)
	}

	return nil
}
