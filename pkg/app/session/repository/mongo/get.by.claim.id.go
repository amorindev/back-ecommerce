package mongo

import (
	"context"
	"errors"

	sessionErr "com.fernando/pkg/app/session/errors"
	"com.fernando/pkg/app/session/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// * establecer el standar sessionErr nombre

func (r *Repository) GetByClaimID(ctx context.Context, id string) (*model.Session, error) {
	var session model.Session

	err := r.Collection.FindOne(ctx, bson.D{{Key: "refresh_token_id", Value: id}}).Decode(&session)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, sessionErr.ErrSessionNotFound
		}
		return nil, err
	}

	// * _id
	sessionObjID, ok := session.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("session mongo repo - session GetByClaimID, failed to parse ID to ObjectID")
	}
	session.ID = sessionObjID.Hex()

	// * user_id
	userObjID, ok := session.UserID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("user mongo repo - user GetByClaimID, failed to parse ID to ObjectID")
	}
	session.UserID = userObjID.Hex()

	return &session, nil
}
