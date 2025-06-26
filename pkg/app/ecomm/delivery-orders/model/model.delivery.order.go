package model

// * Order ID
// En este modelo y en pickup u otro métod va ser primary key es la relacion uno auni
// para postgresql tenermo en cuenta
// * AddresID
// no se si relacionarlo
// falta si se va relacionar ocn adress number con ambos o el usuario tipea desde cero
type DeliveryOrder struct {
	ID        interface{} `json:"id" bson:"_id"`
	OrderID   interface{} `json:"order_id" bson:"order_id"`
    PhoneID   interface{} `json:"phone_id" bson:"phone_id"`
	AddressID interface{} `json:"address_id" bson:"address_id"`
	Reference string      `json:"reference" bson:"reference"`
}

// fecha de entrega estimada

// ? mejor completo y no una relacion conaddressID?
/*
CREATE TABLE delivery_orders (
    order_id UUID PRIMARY KEY REFERENCES orders(id) ON DELETE CASCADE,
    address TEXT NOT NULL,
    city TEXT,
    postal_code TEXT,
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6)
);
*/