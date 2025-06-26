package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/variation-option/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, varOption *model.VariationOption) error {
	id := bson.NewObjectID()
	varOption.ID = id

	variationObjID, err := bson.ObjectIDFromHex(varOption.VariationID.(string))
	if err != nil {
	  return fmt.Errorf("varOption mongo repo - Create err %v", err)
	}
	varOption.VariationID = variationObjID

	_, err = r.Collection.InsertOne(ctx,varOption)
	if err != nil {
		return fmt.Errorf("varOption mongo repo - Create err %v", err)
	}
	// * simpre devolver a stringigual para arreglos -VER EL TEMA DE ARREGLOS MEJOR DESDE EL SERVICO
	// * GET SI EXISTE AHI MISMO RETORNAR EL ERROR, Y NO ESTAR BUSCANDO DENTRO DE CREATE FUNCÍN
	varOption.ID = id.Hex()
	varOption.VariationID = variationObjID.Hex()
	return nil
}


