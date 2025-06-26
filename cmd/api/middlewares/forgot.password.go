package middlewares

// * Solo usar access.token y enviar el token type y validarlo en el handler
/* func ForgotPasswordMiddleware(next http.HandlerFunc) http.HandlerFunc {
	//fmt.Printf("Forgot password midleware ------------\n")
	return func(w http.ResponseWriter, r *http.Request) {
		// obtener la porpose del authmidleware
		purpose, ok := r.Context().Value(PurposeIDKey).(string)
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(message.ErrorMessage{Message: "invalid purpose context"})
			return
		}
		if purpose != "forgot-password" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(message.ErrorMessage{Message: "invalid purpose"})
			return
		}
		next.ServeHTTP(w, r)
	}
} */

/*
¿Y si quiero pasar objetos completos?
También puedes pasar structs completos:

go
Copiar
Editar
ctx = context.WithValue(ctx, UserClaimsKey, c) // `c` puede ser un struct con info del usuario
Y lo recuperas con:

go
Copiar
Editar
if claims, ok := r.Context().Value(UserClaimsKey).(YourClaimsType); ok {
    // usar claims
}
*/
