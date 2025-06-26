package model

type ProductVariation struct {
	ID          interface{} `bson:"_id"`
	ProductID   interface{} `bson:"product_id"`
	VarOptionID interface{} `bson:"var_option"`
}
