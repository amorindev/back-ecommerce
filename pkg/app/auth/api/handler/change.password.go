package handler

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/pkg/app/auth/api/core"
)

// limite de intentos

// ChangePassword cuando el usuario esta logueado
func (h Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	// token válido del login - crear una función de token válido o middleware
	// que se necesita para verificar el token
	w.Header().Set("Content-Type", "application/json")

	var cp core.ChangePasswordRequest

	err := json.NewDecoder(r.Body).Decode(&cp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	defer r.Body.Close()

	// validaciones - logitud del password y caracteres
	if cp.CurrentPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "current password is required"})
		return
	}
	if cp.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "password is required"})
		return
	}
	if cp.ConfirmPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "confirm password is required"})
		return
	}

	// validar fortaleza de la contraseña
	// comparar contraseñas con la base de datos es necesario sin desifrar
	// cyfrar contraseña
	// actuaizar en la base de datos
	// onvaidar tokens/ sesiones

	// servicio
	// - verificar su contraseña actual - neceista email o userID?

	// traducir a request

}

// como manejaría los handlers el mismo middleware?

// token o sesión? - crear sesiones con informacion de la plataforma

//Invalida todos los tokens/sesiones activos excepto el de la solicitud actual, obligando al usuario a iniciar sesión nuevamente en otros dispositivos. excepto a solicitud actual? creo que se debe eliminar y redirigir al login

func GenerateCSRFToken() string {
	token := make([]byte, 32)
	rand.Read(token)

	return base64.URLEncoding.EncodeToString(token)
}

/*
3. ¿Cuándo usar tokens CSRF y cuándo no?
Usa tokens CSRF si:

Tu aplicación utiliza cookies para la autenticación (sesiones tradicionales).
Las cookies de sesión son enviadas automáticamente por el navegador.
No necesitas tokens CSRF si:

Tu aplicación utiliza JWT enviados manualmente en los encabezados (Authorization: Bearer <token>).
Habilitas CORS estricto para aceptar solicitudes solo de orígenes confiables.
*/
