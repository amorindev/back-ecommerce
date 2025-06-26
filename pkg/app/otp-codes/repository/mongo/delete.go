package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Delete(ctx context.Context, id string) error {

	otpObjID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("opt-codes mongo repo - Delete err: %w", err)
	}

	filter := bson.M{"_id": otpObjID}

	//save 4-9 contime aout o cual usar de la transación? otro
	result, err := r.Collection.DeleteOne(ctx, filter)
	// ! son diferentes asi para todos los repos deberia usar result? en todos
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("otp-not-found-id")
	}
	return nil
}
