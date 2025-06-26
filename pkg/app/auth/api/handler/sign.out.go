package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"com.fernando/cmd/api/message"
	"com.fernando/cmd/api/middlewares"
	"com.fernando/internal/claim"
)

// ? usar el middleware de refresh?
func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// ver si se va a validar
	refreshTokenClaims, ok := r.Context().Value(middlewares.RefreshTokenClaimsKey).(*claim.RefreshTokenClaims)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "failed to parse claims - RefreshTokenHandler"})
		return
	}

	// * Validar
	// en este punto la purpose es correcta por que estamos usando un midleware una direrente firma
	// solo para el refreshtoken JWT_REFRESH_STRING ver asi para los demas
	if refreshTokenClaims.TokenType != "refresh-token" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "invalid-purpose RefreshToken"})
		return
	}

	if refreshTokenClaims.ID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "user-not-found-on-claim"})
		return
	}

	// mejor el middleware  que te pase los datos  si no se tiene que vaidar usamos este
	// con todo lo de decode
	/* var req struct {
	Refreshtoken string `json:"refresh_token"`
	}
	defer r.Body.Close()
	*/

	err := h.AuthService.RevokeToken(context.Background(), refreshTokenClaims.ID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "user-not-found-on-claim"})
		return
	}
	// seria 204 no content asi para los demás
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(nil)

	// seisones o tokens?, solo la sesion mobile por ejemplo
	// remove token and refresh mobile and database or revoket = true
}

// ! si se debe vaidar en el mideware porqu igual se borraan en local

/*

    c.JSON(http.StatusOK, gin.H{"message": "Signed out successfully"})
*/

// *  Creo que si expira el refrehs token lo enviamos como error se limpia el localstorage
// * y sign in suppongo
// * solo se usaria para eliminar sessiones, el handler
// * podemos usar go rutina
// * si no hay el token solo no existe el token
// * en que momento el usario se debe quedar con signedin y mostrando el error
// * por que no afecta la consitencia de los datos
// * la unica es que no sea el token válido las demas si salen error o no solo limpiamos el
// * local-storage


/*
1. En el Frontend (Flutter)
Se eliminan los tokens del almacenamiento local (FlutterSecureStorage o shared_preferences).
Esto evita que el usuario pueda seguir haciendo peticiones autenticadas.
Se envía una petición al backend para invalidar el refresh token.
*/
// *Si el backend responde con éxito (204), el logout fue correcto.
// *Si falla (ej: no hay conexión), igual se borran los tokens localmente (pero el refresh token seguirá siendo válido hasta que expire). pero no abri problema igualmenete o mandamos al sign in screen
/*
2. En el Backend
Se recibe el refresh token y se verifica si existe en la base de datos.
Se elimina o marca como inválido (en DB o Redis).
Esto evita que el refresh token pueda usarse para obtener un nuevo access token.
No se hace nada con el access token (sigue siendo válido hasta que expire).
Por eso es importante que los access tokens tengan un tiempo de vida corto (ej: 15-30 minutos).
*/