package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/phones/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetPhonetwoFaSms(ctx context.Context, userID string) (*model.Phone, error) {
	userOID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("GetPhonetwoFaSms: ID inválido: %w", err)
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"user_id": userOID}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "phones", // nombre de la colección de teléfonos
			"localField":   "phone_id",
			"foreignField": "_id",
			"as":           "phone",
		}}},
		{{Key: "$unwind", Value: "$phone"}},
	}

	cursor, err := r.TwoFaSmsColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("GetPhonetwoFaSms: error en agregación: %w", err)
	}
	defer cursor.Close(ctx)

	if !cursor.Next(ctx) {
		return nil, fmt.Errorf("GetPhonetwoFaSms: no se encontró teléfono para userID=%s", userID)
	}

	var result struct {
		Phone model.Phone `bson:"phone"`
	}
	if err := cursor.Decode(&result); err != nil {
		return nil, fmt.Errorf("GetPhonetwoFaSms: error al decodificar resultado: %w", err)
	}

	return &result.Phone, nil
}
