package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/payment/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// * cambia UserID por orderID revisar el flujo que se envi el order id y no user ID
func (r *Repository) Insert(ctx context.Context, orderID string, payment *model.Payment) error {
	id := bson.NewObjectID()
	payment.ID = id

	orderObjID, err := bson.ObjectIDFromHex(orderID)
	if err != nil {
		return fmt.Errorf("payment mongo repo - Insert error: %w", err)
	}
	payment.OrderID = orderObjID

	_, err = r.Collection.InsertOne(ctx, payment)
	if err != nil {
		return fmt.Errorf("user mongo repo - Create error: %w", err)
	}

	payment.ID = id.Hex()
	payment.OrderID = orderObjID.Hex()

	return nil
}

/* func (r *Repository) Insert(ctx context.Context, user *model.User) error {

	aux := user.Roles
	user.Roles = nil


	// lo hago por que primero se crea el user despues los roles si no quedará nil
	user.Roles = aux

	return nil
}
*/
