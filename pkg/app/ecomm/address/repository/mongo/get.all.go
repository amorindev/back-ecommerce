package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/address/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) GetAll(ctx context.Context, userID string) ([]*model.Address, error) {
	oid, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("address - invalid userID: %w", err)
	}
	filter := bson.M{"user_id": oid}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("address - error finding addresses: %w", err)
	}
	defer cursor.Close(ctx)

	var addresses []*model.Address

	// ! hay dos formas de decodificar con cursor Next para convertir el id a string importante
	// ! esta de acuerdo a nuestra logica, pero si no es necesario y vamos a dejar los ids
	// ! de los ids como object id y retornarlos -excelente de momento get all para ser rápidos
	// ! verificar cuando se inserten listas los ids sean obj id insertmany
	if err := cursor.All(ctx, &addresses); err != nil {
		return nil, fmt.Errorf("address - error decoding addresses: %w", err)
	}

	return addresses, nil
}
