package constants

// llamarlo types o enums

type TwoFaMethod string

const (
	MethodAuthApp     TwoFaMethod = "auth_app"
	MethodSms         TwoFaMethod = "sms"
	MethodSecurityKey TwoFaMethod = "security_key"
)

// Para validar valores correctos - en la request
func (m TwoFaMethod) IsValid() bool {
	switch m {
	case MethodAuthApp, MethodSms, MethodSecurityKey:
		return true
	}
	return false
}