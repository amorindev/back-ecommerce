package mongo

import (
	"context"
	"fmt"

	//"com.fernando/pkg/app/ecomm/address/model"
	storeM "com.fernando/pkg/app/ecomm/stores/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetAll2(ctx context.Context) ([]*storeM.Store, error) {
	var stores []*storeM.Store

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("store mongo repo- GetAll: %v", err)
	}
	defer cursor.Close(ctx)

	// Itera a través del cursor y decodifica cada documento
	for cursor.Next(ctx) {
		var store storeM.Store
		if err := cursor.Decode(&store); err != nil {
			return nil, fmt.Errorf("store mongo repo- GetAll: %v", err)
		}
		stores = append(stores, &store)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("store mongo repo- GetAll: %v", err)
	}

	return stores, nil

}

func (r *Repository) GetAll(ctx context.Context) ([]*storeM.Store, error) {
	var stores []*storeM.Store

	pipeline := mongo.Pipeline{
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "addresses"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "store_id"},
				{Key: "as", Value: "address"},
			}},
		},
		{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$address"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			}},
		},
	}

	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("store mongo repo - GetAll aggregation: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var store storeM.Store
		if err := cursor.Decode(&store); err != nil {
			return nil, fmt.Errorf("store mongo repo - GetAll decode: %w", err)
		}

		// Mapear address_docs al campo AddressAgt
		/* raw := cursor.Current
		var address model.Address
		if addrDoc, ok := raw.Lookup("address").DocumentOK(); ok {
			if err := bson.Unmarshal(addrDoc, &address); err == nil {
				store.AddressAgt = &address
			}
		} */

		stores = append(stores, &store)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("store mongo repo - GetAll cursor: %w", err)
	}

	return stores, nil
}
