package model

type PickupOrder struct {
	ID        interface{} `json:"id" bson:"_id"`
	OrderID   interface{} `json:"order_id" bson:"order_id"`
	PhoneID   interface{} `json:"phone_id" bson:"phone_id"`
	AddressID interface{} `json:"address_id" bson:"address_id"`
	StoreID   interface{} `json:"store_id" bson:"store_id"`
	// se puede poner mas datos como fecha de recojo
	// como hacer si el producto esta en otro almacen y quiere recogerlo
	// movimiento de inventario?
}

// * el modelado para postgresql será diferente ver
// * por que mongo si o si va a asignar un id
/* CREATE TABLE pickup_orders (
    order_id UUID PRIMARY KEY REFERENCES orders(id) ON DELETE CASCADE,
    store_id UUID NOT NULL
); */
