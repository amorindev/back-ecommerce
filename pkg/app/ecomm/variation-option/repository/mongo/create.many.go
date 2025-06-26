package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/variation-option/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) CreateMany(ctx context.Context, varOptions []*model.VariationOption) error {
	for _, vOpt := range varOptions {
		id := bson.NewObjectID()
		vOpt.ID = id

		// TODO esto me parece que se hae solo una ves mas arriba varOptions[0].VariationID.(string) 
		varOptObjID, err := bson.ObjectIDFromHex(vOpt.VariationID.(string))
		if err != nil {
			return fmt.Errorf("varOptions mongo repo - CreateMany parse err: %w", err)
		}
		vOpt.VariationID = varOptObjID
	}

	session, err := r.Client.StartSession()
	if err != nil {
		return fmt.Errorf("varOptions mongo repo - CreateMany err: %w", err)
	}

	err = session.StartTransaction()
	if err != nil {
		return fmt.Errorf("varOptions mongo repo - CreateMany err: %w", err)
	}

	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(ctx2 context.Context) error {
		// * dentro de aui se usaría el context 2
		_, err := r.Collection.InsertMany(ctx2, varOptions)
		if err != nil {
			return err
		}

		return session.CommitTransaction(ctx2)
	})
	if err != nil {
		if err := session.AbortTransaction(context.Background()); err != nil {
			return fmt.Errorf("varOptions mongo repo - CreateMany abort transaction err: %w", err)
		}

		return fmt.Errorf("varOptions mongo repo - CreateMany err: %w", err)
	}

	// Dos opciones *[]* no podría ser por que no afecta a la ongitud, en la capa de
	// servicio se debería ver reflejado sin problemas usando []*
	// si no retornar el slice
	return nil
}
