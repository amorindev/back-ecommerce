package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/session/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Create(ctx context.Context, session *model.Session) error {
	objID := bson.NewObjectID()
	session.ID = objID

	// * se que viene string por eso no lo valido
	userObjID, err := bson.ObjectIDFromHex(session.UserID.(string))
	if err != nil {
		return fmt.Errorf("user mongo repo - AssignRoles err 1: %w", err)
	}

	session.UserID = userObjID

	_, err = r.Collection.InsertOne(ctx, session)
	if err != nil {
		return fmt.Errorf("session mongo repo - Create err: %w", err)
	}

	session.ID = objID.Hex()
	session.UserID = userObjID.Hex()

	return nil
}
