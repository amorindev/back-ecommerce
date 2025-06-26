package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/product-config/model"
	productModel "com.fernando/pkg/app/ecomm/product-item/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) InsertMany(ctx context.Context, products []*productModel.ProductItem) error {
	var productConfig []*model.ProductVariation

	for _, product := range products {

		productObjID, err := bson.ObjectIDFromHex(product.ID.(string))
		if err != nil {
			return err
		}

		for _, opt := range product.Options {
			id := bson.NewObjectID()

			optObjID, err := bson.ObjectIDFromHex(opt.VarOptionID.(string))
			if err != nil {
				return err
			}

			pConfig := &model.ProductVariation{
				ID:          id,
				ProductID:   productObjID,
				VarOptionID: optObjID,
			}
			productConfig = append(productConfig, pConfig)
		}
	}

	//if len(relations) > 0 { ? envolver InsertMany en este if ?
	if len(productConfig) <= 0 {
		return nil
		//return errors.New("product-config mongo repo-InsertMany len <=0 ")
	}
	_, err := r.Collection.InsertMany(ctx, productConfig)
	if err != nil {
		return fmt.Errorf("product-config mongo repo - InserMany err: %v", err)
	}

	return nil
}
