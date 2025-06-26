package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/phones/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// ! es importante que context se va a usar
func (t *Transaction) Insert(ctx context.Context, phoneDefaultID string, phone *model.Phone) error {
	session, err := t.Client.StartSession()
	if err != nil {
		return fmt.Errorf("phone mongo tx - SignUpWithOtp err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("phone mongo tx - Cr err: %w", err)
	}

	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(ctx context.Context) error {
		// * no devulve un error si no lo encuentra
		err = t.PhoneRepo.ChangeDefault(ctx, phoneDefaultID, false)
		if err != nil {
			return err
		}

		err = t.PhoneRepo.Insert(ctx, phone)
		if err != nil {
			return err
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err
		}

		return nil
		// return session.CommitTransaction(ctx) o usar esto
	})
	if err != nil {
		if err := session.AbortTransaction(context.Background()); err != nil {
			return fmt.Errorf("phone mongo transaction - Insert AbortTransaction err: %w", err)
		}
		return fmt.Errorf("phone mongo transaction - Insert err: %w", err)
	}

	return nil
}
