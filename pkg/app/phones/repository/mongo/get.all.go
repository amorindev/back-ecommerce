package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/phones/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)


func (r *Repository) GetAll(ctx context.Context, userID string) ([]*model.Phone, error){
	uID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("phone - invalid userID: %w", err)
	}
	filter := bson.M{"user_id": uID}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("phone - error finding addresses: %w", err)
	}
	defer cursor.Close(ctx)

	var phones []*model.Phone

	// ! hay dos formas de decodificar con cursor Next para convertir el id a string importante
	// ! esta de acuerdo a nuestra logica, pero si no es necesario y vamos a dejar los ids
	// ! de los ids como object id y retornarlos -excelente de momento get all para ser rápidos
	// ! verificar cuando se inserten listas los ids sean obj id insertmany
	if err := cursor.All(ctx, &phones); err != nil {
		return nil, fmt.Errorf("phone - error decoding addresses: %w", err)
	}

	return phones, nil
}