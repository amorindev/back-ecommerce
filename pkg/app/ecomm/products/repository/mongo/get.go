package mongo

import (
	"context"
	"fmt"

	"com.fernando/pkg/app/ecomm/products/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Get(ctx context.Context) ([]*model.Product, error) {

	/* pGObjectID, err := bson.ObjectIDFromHex("680853240e1ac5fd7357667b")
	if err != nil {
		return nil, fmt.Errorf("this is a error %w", err)
	} */

	/* pipeline := []bson.M{
		{
			"$match": bson.M{
				"_id": productID, // El ID del Product que estás buscando
			},
		},
		{
			"$lookup": bson.M{
				"from":         "products",
				"localField":   "_id",
				"foreignField": "product_id",
				"as":           "products",
			},
		},
		{
			"$lookup": bson.M{
				"from": "product_variations", // Asumo que esta es la colección intermedia
				"let":  bson.M{"product_ids": "$products._id"},
				"pipeline": []bson.M{
					{
						"$match": bson.M{
							"$expr": bson.M{
								"$in": []interface{}{"$product_id", "$$product_ids"},
							},
						},
					},
					{
						"$lookup": bson.M{
							"from":         "variation_options",
							"localField":   "var_option",
							"foreignField":  "_id",
							"as":           "variation_option",
						},
					},
					{
						"$unwind": "$variation_option",
					},
					{
						"$lookup": bson.M{
							"from":         "variations",
							"localField":   "variation_option.variation_id",
							"foreignField": "_id",
							"as":           "variation",
						},
					},
					{
						"$unwind": "$variation",
					},
				},
				"as": "product_variations",
			},
		},
		{
			"$addFields": bson.M{
				"variations": bson.M{
					"$reduce": bson.M{
						"input": "$product_variations",
						"initialValue": []bson.M{},
						"in": bson.M{
							"$let": bson.M{
								"vars": bson.M{
									"existing": bson.M{
										"$filter": bson.M{
											"input": "$$value",
											"as":    "v",
											"cond": bson.M{
												"$eq": []interface{}{"$$v.name", "$$this.variation.name"},
											},
										},
									},
								},
								"in": bson.M{
									"$cond": bson.M{
										"if": bson.M{"$gt": []interface{}{bson.M{"$size": "$$existing"}, 0}},
										"then": bson.M{
											"$map": bson.M{
												"input": "$$value",
												"as":    "v",
												"in": bson.M{
													"$cond": bson.M{
														"if": bson.M{"$eq": []interface{}{"$$v.name", "$$this.variation.name"}},
														"then": bson.M{
															"name": "$$v.name",
															"values": bson.M{
																"$setUnion": []interface{}{
																	"$$v.values",
																	[]interface{}{"$$this.variation_option.value"},
																},
															},
														},
														"else": "$$v",
													},
												},
											},
										},
										"else": bson.M{
											"$concatArrays": []interface{}{
												"$$value",
												[]bson.M{
													{
														"name": "$$this.variation.name",
														"values": []interface{}{"$$this.variation_option.value"},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			"$project": bson.M{
				"product_variations": 0, // Eliminamos este campo temporal
			},
		},
	} */

	/* pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "products",
				"localField":   "_id",
				"foreignField": "product_id",
				"as":           "products",
			},
		},
		{
			"$lookup": bson.M{
				"from": "product_config", // Colección intermedia
				"let":  bson.M{"product_ids": "$products._id"},
				"pipeline": []bson.M{
					// Filtramos solo las variaciones de los producotos de este grupo
					{
						"$match": bson.M{
							"$expr": bson.M{
								"$in": []interface{}{"$product_id", "$product_ids"},
							},
						},
					},
					// Unimos con las opciones de variación (color, talla) digo "S","M", "ROJO"
					{
						"$lookup": bson.M{
							"from":         "variation_options",
							"localField":   "var_option",
							"foreignField": "_id",
							"as":           "variation_option_info", // es "S" almacenamos temporalmente
						},
					},
					// 3. Descomponemos el array (cada product_variation tiene una sola variation_option)
					{
						"$unwind": "$variation_option_info",
					},
					// 4. Unimos con el tipo de variación (COLOR, SIZE)
					{
						"$lookup": bson.M{
							"from":         "variations",
							"localField":   "variation_option_info.variation_id",
							"foreignField": "_id",
							"as":           "variation_type_info",
						},
					},
					// 5. Descomponemos el array (cada variation_option tiene un solo variation_type)
					{
						"$unwind": "$variation_type_info",
					},
					// 6. Proyectamos para dejar solo los campos que
					{
						"$project": bson.M{
							"product_id":     1,
							"option_value":   "variation_option_info.value",
							"variation_name": "$variation_type_info.name",
						},
					},
				},
			},
			"as": "options",
		},
	} */

	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "product_items",
				"localField":   "_id",
				"foreignField": "product_id",
				"as":           "product_items",
			},
		},
	}
	var products []*model.Product

	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("product mongo repo - Get err:%w", err)
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &products); err != nil {
		return nil, fmt.Errorf("product mongo repo - Get err:%w", err)
	}

	return products, nil

}

/* {
	"$lookup": bson.M{
		"from": "product_config",
		"let":  bson.M{"product_ids": "$products._id"},
		"pipeline": []bson.M{
			{
				"$match": bson.M{
					"$expr": bson.M{"$in": []interface{}{"$product_id", "$$product_ids"}},
				}  ,// Coma eliminada aquí
			},
			{
				"$lookup": bson.M{
					"from":         "variation_options",
					"localField":   "var_option",
					"foreignField": "_id",
					"as":           "option",
				},
			},
			{"$unwind": "$option"},
			{
				"$lookup": bson.M{
					"from":         "variations",
					"localField":   "option.variation_id",
					"foreignField": "_id",
					"as":           "variation",
				},
			},
			{"$unwind": "$variation"},
			{
				"$group": bson.M{
					"_id": "$product_id",
					"options": bson.M{
						"$push": bson.M{
							"name":  "$variation.name",
							"value": "$option.value",
						},
					},
				},
			},
		},
		"as": "products_with_options",
	},
},*/
// ... resto del pipeline
