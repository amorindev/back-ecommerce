package model

type Category struct {
	ID   interface{} `json:"id" bson:"_id"`
	Name *string     `json:"name" bson:"name"`
}
