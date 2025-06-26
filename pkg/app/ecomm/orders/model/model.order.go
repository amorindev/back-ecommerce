package model

import (
	"time"

	"com.fernando/pkg/app/ecomm/delivery-orders/model"
	orderM "com.fernando/pkg/app/ecomm/payment/model"
	pickupM "com.fernando/pkg/app/ecomm/pickup-orders/model"
)

// * me parece que debe incluir el email de quien pago por que la cuenta
// * puede ser de una persona pero  la tarjeta corresponde a otro email ver

// * Al insertar solo te deel id de la direccion para calcuar el costo de envío

// ! que datos se enviaran en el metadata(adjunralo desde mobile) para recibirlo desde el webhook,
// ! hasta ahora el orderID para cambiar el estado a pagado
// enviar el correo que se realizo la compra por tu seguridad
// que pasa si existe un eroor al de al actualizar la compra como pagado que se hace flujo
// en el weebhook? que lo envi nuevamente y le mandamos que se recibio
// VALIDAR QUE CONCIDAN EL ORDER Y EL PAGO EN STRIPE
// crear el order con el payment con stado no pagado
// * no se si va retornar la order al crear por siacaso estoy poniendo el
// * me parece que aqui no se calcula el total sino solo se verifica que coindica precio por cantidad
// * y que coincida con el total,
// * segun cada pais se ve si hay porcentaje de igv u otro
// ! Como los campos de mi sestructura debven ser puntero antes de retornar la entidad se debería validar
// ! que si esten los datos para que si no es asi internal server error, para que no creashee el frontend
// ! asi como establecer en nil passwordHas y password
// TODO si no sale la relaccion lo insertamos todo de una
// ! DeliveryType ver si se va obtener desde la base de datos o es un enum

// ! ver como manejar pickup y delivery entidades de momento todo junto
// * PaymentAgregate - marcar omitempty(setear null en el servicio) o nulos datos sencibles
// * Items - OrderItem noes un agregate de Product{} cuando usar agrgate de momento no lo insertamos
// * Items - depende del modelo si se va a guardar dentro

// * ID - desde la base de datos
// * UserID - desde el token específicamente desde el Subject
type Order struct {
	ID           interface{}          `json:"id" bson:"_id"`
	UserID       interface{}          `json:"user_id" bson:"user_id"`
	Items        []*OrderItem         `json:"order_items" bson:"order_items,omitempty"`
	Total        *float64             `json:"total" bson:"total"`
	DeliveryType string               `json:"delivery_type" bson:"delivery_type"`
	PaymentAgt   *orderM.Payment      `json:"payment" bson:"payment,omitempty"`
	PickupAgt    *pickupM.PickupOrder `json:"pickup" bson:"pickup"`
	DeliveryAgt  *model.DeliveryOrder `json:"delivery" bson:"delivery"`
	CreatedAt    *time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt    *time.Time           `json:"updated_at" bson:"updated_at"`
}

// Name no debe agregarse a la base de datos lo quiero para el response
// y 
type OrderItem struct {
	ID            interface{} `json:"id" bson:"_id"`
	OrderID       interface{} `json:"order_id" bson:"order_id"`
	ProductItemID interface{} `json:"product_item_id" bson:"product_item_id"`
	Quantity      *int        `json:"quantity" bson:"quantity"` // * Guardar en la colleccion intermedia
	Price         *float64    `json:"price" bson:"price"`       // * guardar en la collecion intermedia segunc como sediseñe
	Name          string      `json:"name" bson:"name"` 
}

// ! para eliminar cuenta delete account cascade seria usar transacciones

// * guia
/* type UserRole struct {
	ID     interface{} `json:"-" bson:"_id"`           // not null cascade como hacer con transacciones
	UserID interface{} `json:"user_id" bson:"user_id"` // not null cascade como hacer con transacciones
	RoleID interface{} `json:"role_id" bson:"role_id"` // primary
	// que pasa si tengo mas datos aqui
	// * Creo un servicio user_role ?
} */

/*
CREATE TABLE orders (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    order_status TEXT NOT NULL,
    delivery_type TEXT NOT NULL CHECK (delivery_type IN ('pickup', 'delivery')),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE pickup_orders (
    order_id UUID PRIMARY KEY REFERENCES orders(id) ON DELETE CASCADE,
    store_id UUID NOT NULL
);

CREATE TABLE delivery_orders (
    order_id UUID PRIMARY KEY REFERENCES orders(id) ON DELETE CASCADE,
    address TEXT NOT NULL,
    city TEXT,
    postal_code TEXT,
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6)
);

*/
