package handler

import (
	"net/http"
)

//  NewPassword handler esta relacionado con ForgetPassword handler
// ambos - deben invalidar tokens  o seiones  el usuario puede seleccionar las sesiones
// current user enviar el estado de authenticacion

// Cuando el usuario no inició sesión
func (h Handler) NewPassword(w http.ResponseWriter, r *http.Request) {
	// se envia el token - email -passowe
	// token válido que se envió del email
	// validar que no expire

	// hash new password bycript

	// actualizar en la base de datos

	// invalidar token
}
