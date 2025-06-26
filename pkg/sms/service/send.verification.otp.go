package service

import "fmt"

// ! to := "+51906767533"
//  verificar que el usuario tenga un telefono verificado
func (s *Service) SendVerificationOtp(to string, code string) error {
	// desde el servicio
	msg := fmt.Sprintf("Su código de veridicacion de New Athletic es: %s", code)
	from := "+18316103713"
	// ver este tema, ver que sea dinamico no solo para peru
	completeNumber := fmt.Sprintf("+51%s", to)
	return s.SmsAdp.Send(from, completeNumber, msg)
}
