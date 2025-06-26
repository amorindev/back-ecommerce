package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/variation-option/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// * Seria buscarlo por el nombre
func (r *Repository) ExistOne(ctx context.Context, varOptions []*model.VariationOption) ([]*model.VariationOption, error) {
	// ! crear una funcion aparte ExistOne en el port 
	// * Extraer values atribute para buscar
	values := make([]string, 0, len(varOptions))
	for _, v := range varOptions {
		// ? deepseek esta revisando nulos cuando acerlo al usar?
		if v.Value != nil {
			values = append(values, *v.Value)
		}
	}

	// Consultar MongoDB
	filter := bson.M{"value": bson.M{"$in": values}}
	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error finding variation options: %w", err)
	}
	defer cursor.Close(ctx)

	// Obtener resultados existentes, 
	// ? debería ser puntero ?
	var existingOptions []*model.VariationOption
	if err := cursor.All(ctx, &existingOptions); err != nil {
		return nil, fmt.Errorf("error decoding variation options: %v", err)
	}

	// * otra forma seria si existingOptions > 0 retornar false

	return existingOptions, nil
}
