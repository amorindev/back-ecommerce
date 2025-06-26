package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/permission/model"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) InsertMany(ctx context.Context, permissions []*model.Permission) error {
	for _, p := range permissions {
		id := bson.NewObjectID()
		p.ID = id
	}

	session, err := r.Client.StartSession()
	if err != nil {
		return fmt.Errorf("permissions mongo repo - InsertMany err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("permissions mongo repo - InsertMany err: %w", err)
	}

	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(ctx2 context.Context) error {
		// * dentro de aui se usaría el context 2 ver si es necesario renombralo ctx2
		_, err := r.Collection.InsertMany(ctx2, permissions)
		if err != nil {
			return err
		}

		return session.CommitTransaction(ctx2)
	})
	if err != nil {
		if err := session.AbortTransaction(context.Background()); err != nil {
			return fmt.Errorf("permissions mongo repo - InsertMany abort transaction err: %w", err)
		}

		return fmt.Errorf("permissions mongo repo - InsertMany err: %w", err)
	}
	return nil

}

