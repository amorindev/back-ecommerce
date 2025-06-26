package mongo

import (
	"context"
	"errors"

	authErr "com.fernando/pkg/app/auth/errors"
	"com.fernando/pkg/app/auth/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// ! esta cauando errro por que el estoy pasando el id del user
// ! y se necesita del auth por eso me da not func ver si se va udar el provider
// ! para solo dejar user ver el handler y servicios que sean el correcto getuSer
// ! y no get auth
func (r *Repository) Get(ctx context.Context, id string) (*model.Auth, error) {

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("auth mongo repo - Get invalid type ID")
	}
	var auth model.Auth
	// * probar con el ID
	// ! ahora es mediante en UserID
	err = r.Collection.FindOne(ctx, bson.D{{Key: "_id", Value: objID}}).Decode(&auth)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, authErr.ErrAuthNotFound
		}
		return nil, err
	}

	// ! cambiar a de idStr to objID
	objID, ok := auth.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("auth mongo repo - Get falied to parse ID to string")
	}
	auth.ID = objID.Hex()

	return &auth, nil
}
