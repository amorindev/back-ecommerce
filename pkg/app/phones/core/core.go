package core

import "time"

// IsDefault opcional pero lo mantedremos al crear como default
// verified false
// se tendría que crear otra tabla cyty o pais
// y ahis agregar igv de momento será opcional countryCode
// y ademas como quedaría la relacion
type CreatePhoneReq struct {
	Number         *string    `json:"number" bson:"number"`
	CountryCode    *string    `json:"country_code" bson:"country_code"`
	CountryIsoCode *string    `json:"country_iso_code" bson:"country_iso_code"`
	IsDefault      bool       `json:"is_default" bson:"is_default"`
	CreatedAt      *time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" bson:"updated_at"`
}
