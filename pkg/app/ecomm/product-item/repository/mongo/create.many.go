package mongo

import (
	"context"
	"errors"
	"fmt"

	"com.fernando/pkg/app/ecomm/product-item/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) CreateMany(ctx context.Context, products []*model.ProductItem) error {
	productObjID, err := bson.ObjectIDFromHex(products[0].ProductID.(string))
	if err != nil {
		return fmt.Errorf("varOptions mongo repo - CreateMany parse err: %w", err)
	}

	for _, product := range products {
		id := bson.NewObjectID()
		product.ID = id
		product.ProductID = productObjID
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
		_, err := r.Collection.InsertMany(ctx2, products)
		if err != nil {
			return err
		}

		return session.CommitTransaction(ctx2)
	})

	if err != nil {
		if err := session.AbortTransaction(context.Background()); err != nil {
			return fmt.Errorf("CreateMany mongo repo - CreateMany abort transaction err: %w", err)
		}

		return fmt.Errorf("CreateMany mongo repo - CreateMany err: %w", err)
	}

	// si desde otro lugar queremos parsear a bson desde lo usemos podemos varificar si es bson
	// y yano hacer lo siguiente, pero para matener coherencia simpre salir con string
	//panic: interface conversion: interface {} is bson.ObjectID, not string
	for _, p := range products {
		idObj, ok := p.ID.(bson.ObjectID)
		if !ok {
			return errors.New("CreateMany mongo repo - CreateMany parse to objID")
		}
		p.ID = idObj.Hex()
	}

	return nil
}

/* func (r *Repository) CreateMany2(ctx context.Context, varOptions []*model.VariationOption) error {
	for _, vOpt := range varOptions {
		varOptObjID, err := bson.ObjectIDFromHex(vOpt.VariationID.(string))
		if err != nil {
			return fmt.Errorf("varOptions mongo repo - CreateMany parse err: %w", err)
		}
		vOpt.VariationID = varOptObjID
	}

	// Dos opciones *[]* no podría ser por que no afecta a la ongitud, en la capa de
	// servicio se debería ver reflejado sin problemas usando []*
	// si no retornar el slice
	return nil
}
*/
