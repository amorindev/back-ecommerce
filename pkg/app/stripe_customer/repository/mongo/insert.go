package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/stripe_customer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, pProvider *model.StripeCustomer) error {
	id := bson.NewObjectID()
	pProvider.ID = id

	userObjID, err := bson.ObjectIDFromHex(pProvider.UserID.(string))
	if err != nil {
		return fmt.Errorf("stripe_customer mongo err: %w", err)
	}
	customerObjID, err := bson.ObjectIDFromHex(pProvider.CustomerID.(string))
	if err != nil {
		return fmt.Errorf("stripe_customer mongo err: %w", err)
	}

	pProvider.UserID = userObjID
	pProvider.CustomerID = customerObjID

	_, err = r.Collection.InsertOne(context.Background(), pProvider)
	if err != nil {
		return fmt.Errorf("stripe_customer mongo repo: Insert err: %w", err)
	}
	pProvider.ID = id.Hex()
	pProvider.UserID = userObjID.Hex()
	pProvider.CustomerID = customerObjID.Hex()
	//!if result.InsertedID == nil ?  todos
	return nil
}
