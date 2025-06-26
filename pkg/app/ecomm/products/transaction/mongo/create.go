package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/products/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (t *Transaction) Create(ctx context.Context, product *model.Product) error {
	session, err := t.Client.StartSession()
	if err != nil {
		return fmt.Errorf("product mongo tx - Create err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("product mongo tx - Create err: %w", err)
	}

	defer session.EndSession(ctx)

	// * necesitamos context en Create como diferenciar ambos
	err = mongo.WithSession(ctx, session, func(ctx2 context.Context) error {
		err = t.ProductRepo.Insert(ctx2, product)
		if err != nil {
			return err
		}

		err = t.ProductItemRepo.CreateMany(ctx2, product.ProductItems)
		if err != nil {
			return err
		}

		// ! Aqui hay un problema no me deberia crear lo demas y en la base de datos si me aparece
		err = t.ProductConfigRepo.InsertMany(context.Background(), product.ProductItems)
		if err != nil {
			return err
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		if err := session.AbortTransaction(context.Background()); err != nil {
			return fmt.Errorf("product mongo transaction - Create AbortTransaction err: %w", err)
		}
		return fmt.Errorf("product mongo transaction - Create err: %w", err)
	}

	return nil

}
