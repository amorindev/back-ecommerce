package model

import (
	"time"
)

// * si lo relaciono con auth
// * lo que pasa que esta relacionada con auth de provider password
// Purpose 'verify_email', 'reset_password', 'sign-up', 'update-email', 'delete-account'
type OtpCodes struct {
	ID        interface{} `bson:"_id"`
	UserID    interface{} `bson:"auth_id"`
	OptCode   *string     `bson:"otp_code"`
	Purpose   *string     `bson:"purpose"` 
	ExpiresAt *time.Time  `bson:"expires_at"`
	Used      *bool       `bson:"used"`
	CreatedAt *time.Time  `bson:"created_at"`
}

//🛠️ Alternativa (menos granular pero válida)
//Si solo te interesa OTP por usuario, no importa el método de autenticación:
//Puedes relacionarlo directamente con user_id en vez de auth_id.
