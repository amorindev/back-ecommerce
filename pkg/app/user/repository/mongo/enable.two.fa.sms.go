package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/user/constants"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) EnableTwoFaSms(ctx context.Context, userID string, twoFaMethod constants.TwoFaMethod) error {
	//fmt.Printf("User ID repo: %s\n", userID)
	userOID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("user mongo repo EnableTwoFaSms failed: %w", err)
	}
	filter := bson.M{
		// o id que es 1 a 1 asegura te de twoFaMethod.ID = userID ver
		"_id": userOID,
	}
	update := bson.M{
		"$set": bson.M{
			"is_2fa_enabled": true,
			"two_fa_method":  twoFaMethod, // opcional: si quieres guardar el método elegido
		},
	}
	result, err := r.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		// seria not found pero no se confundiría con el delete cuando no lo encuentra
		// * me parece que no go senior en cada cada cambia el mensaje de error
		return fmt.Errorf("EnableTwoFaSms: no se encontró ningún documento para actualizar con userID=%s", userID)

		//return errors.New("no se encontró ningún documento para actualizar")
	}
	// result.MatchedCount estudiar los demas
	//fmt.Printf("Documentos modificados %v\n", result.ModifiedCount)

	return nil
}

/* func (r *Repository) EnableTwúoFaSms(ctx context.Context, userID string, twoFaMethod constants.TwoFaMethod) error {
	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
	}

	if result.MatchedCount == 0 {
		return errors.New("EnableTwoFaSms: no document matched the filter criteria")
	}

	if result.ModifiedCount == 0 {
		// Ya estaba en el estado deseado
		return nil
	}

	return nil
}
*/
