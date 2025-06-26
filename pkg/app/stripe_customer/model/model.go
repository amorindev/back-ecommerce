package model

import "time"

// en postgrsql no se necesitaría el ID
type StripeCustomer struct {
	ID         interface{} `bson:"_id"`
	UserID     interface{} `bson:"user_id"`
	CustomerID interface{} `bson:"customer_id"`
	CreatedAt  *time.Time  `bson:"created_at"`
}
