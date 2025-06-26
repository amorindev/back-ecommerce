package mongo

import (
	"context"
	"errors"
	"fmt"

	sCustomerErr "com.fernando/pkg/app/stripe_customer/errors"
	"com.fernando/pkg/app/stripe_customer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetByUserID(ctx context.Context, userID string) (*model.StripeCustomer, error) {
	userObjID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("stripe_customer mongo repo - GetByUserID err d: %w", err)
	}

	// ver si es necesario parsearlo a obj o sino usar el userID
	filter := bson.M{"user_id": userObjID}
	
    var customer model.StripeCustomer
	// no deberia haber mas de uno
    err = r.Collection.FindOne(ctx, filter).Decode(&customer)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, sCustomerErr.ErrStripeCustomerNotFound 
        }
        return nil, err 
    }

	customer.UserID = userObjID.Hex()

	objID, ok := customer.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("stripe_customer mongo repo - GetByUserID, failed to convert authID to ObjectID")
	}
	customer.ID = objID.Hex()
    return &customer, nil
}

