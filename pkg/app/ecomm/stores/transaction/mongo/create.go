package mongo

import (
	"context"
	"fmt"

	addressM "com.fernando/pkg/app/ecomm/address/model"
	storeM "com.fernando/pkg/app/ecomm/stores/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (t *Transaction) Create(ctx context.Context, store *storeM.Store, address *addressM.Address) error {
	session, err := t.Client.StartSession()
	if err != nil {
		return fmt.Errorf("store mongo tx - Create err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("store mongo tx - Create err: %w", err)
	}

	// * La ves pasada verificaste que aunque salia un error estaba guardandose en la base de
	// * datos puede ser por que estas pasando context background u otro

	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(ctx context.Context) error {
		err = t.StoreRepo.Insert(ctx, store)
		if err != nil {
			return err
		}

		address.StoreID = store.ID // usar el .(string)?,
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
			return fmt.Errorf("store mongo transaction - Create AbortTransaction err: %w", err)
		}
		return fmt.Errorf("store mongo transaction - Create err: %w", err)
	}
	return nil
}
