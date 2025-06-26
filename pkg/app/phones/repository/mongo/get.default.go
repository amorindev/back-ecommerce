package mongo

import (
	"context"
	"errors"

	phoneErr "com.fernando/pkg/app/phones/errors"
	"com.fernando/pkg/app/phones/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


func (r *Repository) GetDefault(ctx context.Context) (*model.Phone, error){
	var phone model.Phone

	err := r.Collection.FindOne(ctx, bson.D{{Key: "is_default", Value: true}}).Decode(&phone)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, phoneErr.ErrPhoneNotFound
		}
		// verificar que todos retornen esto
		return nil, err
	}

	objID, ok := phone.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("phone mongo repo - GetDefault, failed to parse ID to ObjectID")
	}
	phone.ID = objID.Hex()

	userObjID, ok := phone.UserID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("address mongo repo - GetDefault, failed to parse ID to ObjectID")
	}
	phone.UserID = userObjID.Hex()
	return &phone,nil
}