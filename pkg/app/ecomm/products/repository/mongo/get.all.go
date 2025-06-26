package mongo

import (
	"context"

	"com.fernando/pkg/app/ecomm/products/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// * 1 * 10 - 10 = 0
// * 2 * 10 - 10 = 10
// * 3 * 10 - 10 = 20

func (r *Repository) GetAll(ctx context.Context, limit int, page int) ([]*model.Product, error) {
	

	// filatra por productos activos
	l := int64(limit)
	skip := int64(page*limit - limit)
	fOptions := options.Find().SetLimit(l).SetSkip(skip)
	cursor, err := r.Collection.Find(ctx, bson.D{}, fOptions)
	if err != nil {
		return nil, err
	}
	// TODO aqui si cancelar con la request por que es un get, o pasarun context con time
	defer cursor.Close(ctx)

	// ! COmo hacer la consulta entro todos las colecciones y en el servicio comoagregar las
	// ! variaciones

	var products []*model.Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}


/*
If you need to set more options, you can chain them like:


findOptions := options.Find().
    SetLimit(l).
    SetSkip(skip).
    SetSort(bson.D{{"field", 1}})
*/

// ! filtrar por active

/* func (r *ProductRepository) GetAll(ctx context.Context, limit int, page int) ([]Product, error) {
	l := int64(limit)
	skip := int64(page*limit - limit)
	
	findOptions := options.Find().
		SetLimit(l).
		SetSkip(skip).
		SetSort(bson.D{{"name", 1}}) // Ordenar por nombre
	
	// Filtro para productos activos
	filter := bson.D{{"active", true}}
	
	curr, err := r.Collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("error al buscar productos: %v", err)
	}
	defer curr.Close(ctx)
	
	var products []Product
	if err := curr.All(ctx, &products); err != nil {
		return nil, fmt.Errorf("error al decodificar productos: %v", err)
	}
	
	return products, nil
} */