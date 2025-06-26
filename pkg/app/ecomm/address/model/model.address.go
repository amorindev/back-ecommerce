package model

import "time"

// * Label
// nombre para mostrar "Casa" "Oficina"
type Address struct {
	ID          interface{} `json:"id" bson:"_id"`
	UserID      interface{} `json:"user_id" bson:"user_id"`
	StoreID     interface{} `json:"store_id" bson:"store_id"`
	Label       *string     `json:"label" bson:"label"`
	AddressLine string      `bson:"address_line" json:"address_line"`                 // Dirección formateada
	City        string      `bson:"city,omitempty" json:"city,omitempty"`             // Ciudad (opcional)
	State       string      `bson:"state,omitempty" json:"state,omitempty"`           // Estado o provincia (opcional)
	Country     string      `bson:"country" json:"country"`                           // País
	PostalCode  string      `bson:"postal_code" json:"postal_code"`                   // Código postal
	Latitude    float64     `bson:"latitude" json:"latitude"`                         // Latitud (DECIMAL(10,8))
	Longitude   float64     `bson:"longitude" json:"longitude"`                       // Longitud (DECIMAL(11,8))
	IsDefault   bool        `bson:"is_default" json:"is_default"`                     // Dirección principal
	CreatedAt   *time.Time  `bson:"created_at,omitempty" json:"created_at,omitempty"` // Fecha de creación
	UpdatedAt   *time.Time  `bson:"updated_at,omitempty" json:"updated_at,omitempty"` // Fecha de actualización
}

// ! como validar que solo exista un default

// que pasa si es tipo treello o como youtube donde
// el orden de los card si importa y desde la ui puedes ordenar arasstrando hacia arriba
// entonces address debería tener un campo Position y en la ui ordenarlo por position
/*
CREATE TABLE addresses (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    label VARCHAR,
    address_line TEXT,
    latitude DECIMAL(10,8),
    longitude DECIMAL(11,8),
    city VARCHAR,
    state VARCHAR,
    country VARCHAR,
    postal_code VARCHAR,
    is_default BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
*/

/*
{
  "user_id": "uuid-del-usuario",
  "label": "Casa",
  "address_line": "Av. Siempre Viva 742, Springfield",
  "latitude": -12.04318,
  "longitude": -77.02824,
  "city": "Springfield",
  "state": "Illinois",
  "country": "USA",
  "postal_code": "12345",
  "is_default": true
}
*/
