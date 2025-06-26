package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/orders/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// * ver la relacion con payments por que no es uno a muchos ver db yotube
func (r *Repository) GetAll2(ctx context.Context, userID string) ([]*model.Order, error) {
	oID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("order - invalid userID: %w", err)
	}
	filter := bson.M{"user_id": oID}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("order - error finding addresses: %w", err)
	}
	defer cursor.Close(ctx)

	var orders []*model.Order
	// ! hay dos formas de decodificar con cursor Next para convertir el id a string importante
	// ! esta de acuerdo a nuestra logica, pero si no es necesario y vamos a dejar los ids
	// ! de los ids como object id y retornarlos -excelente de momento get all para ser rápidos
	// ! verificar cuando se inserten listas los ids sean obj id insertmany

	if err := cursor.All(ctx, &orders); err != nil {
		return nil, fmt.Errorf("orders - error decoding addresses: %w", err)
	}
	return orders, nil
}

func (r *Repository) GetAll(ctx context.Context, userID string) ([]*model.Order, error) {
	oID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid userID: %w", err)
	}

	/* pipeline:= mongo.Pipeline{
	bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: oID}}}},
	bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "payments"},         // colección de destino
		{Key: "localField", Value: "_id"},        // campo en Order
		{Key: "foreignField", Value: "order_id"}, // campo en Payment
		{Key: "as", Value: "payment"},
	}}},

	bson.D{{Key: "$unwind", Value: bson.D{
		{Key: "path", Value: "$payment"},
		{Key: "preserveNullAndEmptyArrays", Value: true}, // por si no hay pago aún
	}}}, */

	// ! Join con productos - se relaciona con product item despues con productos
	// !de momento solo
	/* {{
		Key: "$lookup", Value: bson.M{
			"from":         "products",
			"localField":   "order_items.product_id",
			"foreignField": "_id",
			"as":           "product",
		},
	}},
	{{Key: "$unwind", Value: "$product"}},

	// Agrupar productos por orden
	{{
		Key: "$group", Value: bson.M{
			"_id":        "$_id",
			"user_id":    bson.M{"$first": "$user_id"},
			"created_at": bson.M{"$first": "$created_at"},
			"products": bson.M{"$push": bson.M{
				"product":  "$product",
				"price":    "$order_items.price",
				"quantity": "$order_items.quantity",
			}},
		},
	}}, */
	/* } */
	/* pipeline := mongo.Pipeline{
		// 1. Filtrar por usuario
		bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: oID}}}},

		// 2. Join con pagos
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "payments"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "order_id"},
			{Key: "as", Value: "payment"},
		}}},
		bson.D{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$payment"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}}},

		// 3. Join con order_product
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "order_product"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "order_id"},
			{Key: "as", Value: "order_items"},
		}}},

		// 4. Unwind de order_items para poder hacer el lookup individual
		bson.D{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$order_items"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}}},

		// 5. Lookup a products usando $expr y $toObjectId
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "products"},
			{Key: "let", Value: bson.D{
				{Key: "prod_id", Value: "$order_items.product_item_id"},
			}},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{
					{Key: "$expr", Value: bson.D{
						{Key: "$eq", Value: bson.A{"$_id", bson.D{{Key: "$toObjectId", Value: "$$prod_id"}}}},
					}},
				}}},
				bson.D{{Key: "$project", Value: bson.D{
					{Key: "name", Value: 1},
				}}},
			}},
			{Key: "as", Value: "product_info"},
		}}},

		// 6. Unwind del resultado del lookup
		bson.D{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$product_info"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}}},

		// 7. Agregar nombre al order_item
		bson.D{{Key: "$addFields", Value: bson.D{
			{Key: "order_items.name", Value: "$product_info.name"},
		}}},

		// 8. Reagrupar los order_items por orden
		bson.D{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$_id"},
			{Key: "user_id", Value: bson.D{{Key: "$first", Value: "$user_id"}}},
			{Key: "created_at", Value: bson.D{{Key: "$first", Value: "$created_at"}}},
			{Key: "total", Value: bson.D{{Key: "$first", Value: "$total"}}},
			{Key: "delivery_type", Value: bson.D{{Key: "$first", Value: "$delivery_type"}}},
			{Key: "payment", Value: bson.D{{Key: "$first", Value: "$payment"}}},
			{Key: "order_items", Value: bson.D{{Key: "$push", Value: "$order_items"}}},
		}}},
	} */
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: oID}}}},

		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "payments"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "order_id"},
			{Key: "as", Value: "payment"},
		}}},
		bson.D{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$payment"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}}},

		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "order_product"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "order_id"},
			{Key: "as", Value: "order_items"},
		}}},

		bson.D{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$order_items"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}}},

		// Join con product_items para obtener product_id
		// ! probar arriba por que lo que estaba mal es las referencias losnombres de
		// ! las collecionesy campos
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "product_items"},
			{Key: "let", Value: bson.D{{Key: "pi_id", Value: "$order_items.product_item_id"}}},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{
					{Key: "$expr", Value: bson.D{
						{Key: "$eq", Value: bson.A{"$_id", bson.D{{Key: "$toObjectId", Value: "$$pi_id"}}}},
					}},
				}}},
				bson.D{{Key: "$project", Value: bson.D{
					{Key: "product_id", Value: 1},
				}}},
			}},
			{Key: "as", Value: "product_item_data"},
		}}},
		bson.D{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$product_item_data"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}}},

		// Join con products para obtener el nombre del producto
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "products"},
			{Key: "let", Value: bson.D{{Key: "prod_id", Value: "$product_item_data.product_id"}}},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{
					{Key: "$expr", Value: bson.D{
						{Key: "$eq", Value: bson.A{"$_id", bson.D{{Key: "$toObjectId", Value: "$$prod_id"}}}},
					}},
				}}},
				bson.D{{Key: "$project", Value: bson.D{
					{Key: "name", Value: 1},
				}}},
			}},
			{Key: "as", Value: "product_info"},
		}}},
		bson.D{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$product_info"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}}},

		bson.D{{Key: "$addFields", Value: bson.D{
			{Key: "order_items.name", Value: "$product_info.name"},
		}}},

		// Regroup por orden
		bson.D{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$_id"},
			{Key: "user_id", Value: bson.D{{Key: "$first", Value: "$user_id"}}},
			{Key: "created_at", Value: bson.D{{Key: "$first", Value: "$created_at"}}},
			{Key: "updated_at", Value: bson.D{{Key: "$first", Value: "$updated_at"}}},
			{Key: "total", Value: bson.D{{Key: "$first", Value: "$total"}}},
			{Key: "delivery_type", Value: bson.D{{Key: "$first", Value: "$delivery_type"}}},
			{Key: "payment", Value: bson.D{{Key: "$first", Value: "$payment"}}},
			{Key: "order_items", Value: bson.D{{Key: "$push", Value: "$order_items"}}},
		}}},
		bson.D{{Key: "$sort", Value: bson.D{{Key: "created_at", Value: -1}}}},
	}

	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregate error: %w", err)
	}
	defer cursor.Close(ctx)

	var results []*model.Order
	if err := cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("error decoding result: %w", err)
	}
	return results, nil
}
