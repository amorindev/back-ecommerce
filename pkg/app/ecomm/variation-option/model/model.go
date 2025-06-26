package model

type VariationOption struct {
	ID          interface{} `bson:"_id"` // ! De momento no se como se verá en el json
	Value       *string     `bson:"value"`
	VariationID interface{} `bson:"variation_id"`
}
