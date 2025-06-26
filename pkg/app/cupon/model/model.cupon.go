package model

import "time"

type DiscountType string

const Percent DiscountType = "percent"
const Fixed DiscountType = "fixed"

// * MinPurcharses
// este me parece que es del producto,
// descuento del producto por que esta detereorado por ejemplo, o esta mucho tiempo en el almacén
// descuento por la compra dechas festivas o tarjeta
// cupon - apps de tipo saas ver
// como sincronizar cupones descuentode producto, y descuento de factura por CRM
// * UserID
// NULL si es público; FK a users si es personal
type Cupon struct {
	ID            interface{}   `json:"id" bson:"_id"`
	Code          *string       `json:"code" bson:"code"` // unique not null -BIENVENIDO10
	Description   *string       `json:"description" bson:"description"`
	DiscountType  *DiscountType `json:"discount_type" bson:"discount_type"`
	DiscountValue *float32      `json:"discount_value" bson:"discount_value"`
	MaxUses       *int          `json:"max_uses" bson:"max_uses"`
	Uses          *int          `json:"uses" bson:"uses"`
	MinPurcharses *float32      `json:"min_purcharses" bson:"min_purcharses"` // uno dominio
	StartDate     *time.Time    `json:"start_date" bson:"start_date"`
	EndDate       *time.Time    `json:"end_date" bson:"ent_date"`
	UserID        *time.Time    `json:"user_id" bson:"user_id"`
	IsActive      *bool         `json:"is_active" bson:"is_active"`
	ExpiresAt     *time.Time    `json:"expires_at" bson:"expires_at"`
	CreatedAt     *time.Time    `json:"created_at" bson:"created_at"`
	UpdatedAt     *time.Time    `json:"updated_at" bson:"updated_at"`
}
