package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/address/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (t *Transaction) Insert(ctx context.Context, addressDefaultID string, address *model.Address) error {
	session, err := t.Client.StartSession()
	if err != nil {
		return fmt.Errorf("address mongo tx - Insert err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("address mongo tx - Insert err: %w", err)
	}

	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(ctx context.Context) error {
		// * no devulve un error si no lo encuentra
		err = t.AddressRepo.ChangeDefault(ctx, addressDefaultID, false)
		if err != nil {
			return err
		}

		err = t.AddressRepo.Insert(ctx, address)
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
			return fmt.Errorf("address mongo transaction - Insert AbortTransaction err: %w", err)
		}
		return fmt.Errorf("address mongo transaction - Insert err: %w", err)
	}

	return nil
}
