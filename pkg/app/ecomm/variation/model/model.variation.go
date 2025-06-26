package model

type Variation struct {
	ID   interface{} `json:"id" bson:"_id"`
	Name *string     `json:"name" bson:"name"`
}
