package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)
// RemoveRoleFromUser gramaticalmete
func (r *Repository) RemoveRolesToUser(ctx context.Context, userID string, roleID string) error {
	userObjID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("role mongo repo - RemoveRolesToUser err: %w", err)
	}

	roleObjID, err := bson.ObjectIDFromHex(roleID)
	if err != nil {
		return fmt.Errorf("role mongo repo - RemoveRolesToUser err: %w", err)
	}

	result, err := r.UserRoleColl.DeleteOne(ctx, bson.M{"user_id": userObjID, "role_id": roleObjID})
	if err != nil {
		return fmt.Errorf("role mongo repo - RemoveRolesToUser err: %w", err)
	}

	if result.DeletedCount == 0 {
		return errors.New("role-not-found-for-user")
	}

	return nil
}
