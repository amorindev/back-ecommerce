package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/category/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) GetAll(ctx context.Context) ([]*model.Category, error) {
	var categories []*model.Category

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("category mongo repo- GetAll: %v", err)
	}
	defer cursor.Close(ctx)

	// Itera a través del cursor y decodifica cada documento
	for cursor.Next(ctx) {
		var category model.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, fmt.Errorf("category mongo repo- GetAll: %v", err)
		}
		categories = append(categories, &category)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("category mongo repo- GetAll: %v", err)
	}

	return categories, nil
}
