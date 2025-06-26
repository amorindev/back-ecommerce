package core

import "time"

// en ves de crear ciudad provincia y distrito crear con recurdividad desde la tabla
// ver primero como te lo da google
// * Address line
// Dirección formateada  (ej: "Av. Siempre Viva 742, Springfield")
// * Label // (ej: "Casa", "Oficina")
// * Latitude // Latitud (DECIMAL(10,8))
// * // Longitud (DECIMAL(11,8))
// * City // Ciudad (opcional)
// * state // Estado o provincia (opcional)
// * country
// TODO para todos los que usan is default marcar por defecto si es el primero en agregar nada mas
// que pasa si es store falta el store id ver flujos
type CreateAddressReq struct {
	Label       *string    `json:"label" `
	AddressLine string     `json:"address_line"`
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	City        string     `json:"city,omitempty"`
	State       string     `json:"state,omitempty"`
	Country     string     `json:"country"`
	PostalCode  string     `json:"postal_code"`
	IsDefault   bool       `json:"is_default"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
