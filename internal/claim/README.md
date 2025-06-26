<!-- * äsar la purpose al handler mediante el context -->
<!-- * de momento solo usamos el id en el refreshtoken para eliminarlo de la base de datos -->
<!-- * cuanodo se retorna el expiresin en sign in? -->

<!-- * si se puede cambiar el id de refresh token de  momento buscarlo por -->
<!-- * refresh_token_id que es uuid de tipo string -->

no necesariamente se necesita un repo
permite pasarle los envs

type TokenService struct {
jwtSecret string
accessTokenExp time.Duration
refreshTokenExp time.Duration
}
func NewTokenService(secret string, accessExp, refreshExp time.Duration) \*TokenService {
return &TokenService{
jwtSecret: secret,
accessTokenExp: accessExp,
refreshTokenExp: refreshExp,
}
}
crear el token con el id del auth o email del usaurio?,

Repositorio (SessionRepository): Se encarga de interactuar con la base de datos para guardar y recuperar sesiones (refresh tokens).
Servicio de Sesión (SessionService): Usa SessionRepository y TokenService para manejar las sesiones y tokens.

Servicio de Token (TokenService): Responsable de generar, validar y refrescar tokens.
Servicio de Autenticación (SignInService): Usa TokenService y SessionService para autenticar usuarios y manejar el inicio de sesión.

debe ir en internal? token pckge definir eso los clientes de resend mongo

<!-- ! token debería estar como servicio dentro de auth?-->

es un tipo de servicio sin base de datos

// \* Token como servicio
// ver microblog o grpc con flutter blog, matenerlo asi o como servicio

//baso en cockies ecomm o goinventory o jwt

<!-- * si quieres validar la purpose lo pasarías  por el context al handler para validarlo -->

## tres tipos de reclamaciones registradas públicas y privadas

# Registradas

- iss : emisor
- asunto : sub
- audiencia : aud
- tiempo de vencimiento : exp
- emitido en : iat

/\* c2 := claim.NewCustomJWT(user.ID, conf.Issuer, userRoles, u.RemenberMe)

signingString, refreshString, err := c2.GetToken(conf.SigningString, conf.RefreshString)
if err != nil {
return "", "", err
}

err = s.TokenService.Create(refreshString, c2.Refresh.ExpiresAt.Time, user.ID)
if err != nil {
return "", "", err
} \*/

```
accessClaims, err := GetAccessTokenFromJWT(accessToken, "my-secret-key")
if err != nil {
    log.Fatal("Invalid access token:", err)
}
fmt.Println("User ID:", accessClaims.UserID)

refreshClaims, err := GetRefreshTokenFromJWT(refreshToken, "my-secret-key")
if err != nil {
    log.Fatal("Invalid refresh token:", err)
}
fmt.Println("User ID:", refreshClaims.UserID)
```

## is valid?
# Firma incorrecta:
- La firma del token no coincide con la calculada usando la clave secreta.
  Token expirado:
- Si el reclamo exp (fecha de expiración) indica que el token ya no es válido.
  Token manipulado:
- Si alguno de los segmentos del token fue modificado después de ser firmado.
Algoritmo de firma desconocido:
- Si el algoritmo especificado en el encabezado no coincide con los algoritmos esperados.
Este campo es útil para evitar que tokens no válidos sean utilizados en operaciones críticas, como la autenticación o la autorización.

refresh_tokens pueden tener referencia directa a user_id (y opcionalmente también al auth_provider_id si quieres saber con qué se logueó el usuario en esa sesión).


🟡 Alternativa: Relacionar user con session
Si tu sistema usa el concepto más general de "sessiones", como en algunas arquitecturas web, puedes reemplazar refresh_tokens por sessions. Ejemplo:
user (1) --- (∞) sessions
sessions --- (1) auth_provider (opcional)
Pero al final suele terminar siendo lo mismo que refresh_tokens, solo con diferente nombre.

Asegúrate de hashear los refresh tokens antes de almacenarlos
Considera agregar campos como IP, user-agent y ubicación para seguimiento de sesiones
Indexa los campos de búsqueda frecuente (user_id, token, expires_at)
Considera un TTL (time-to-live) en la base de datos para limpieza automática


Punteros en golang
1. para no guardar valores que se genera por defecto en go "" 0 si no null (standar)
2. para la api no muestre valores "", o 0 (standar)
3. ahorá como validarlos == nil
4. cual es el flujo, cuando desreferenciar (ver el post de structuras de golang)



cuando crear índices y en que tablas o collecciones

<!-- ! Rotación de tokens: Invalida refresh tokens después de su uso. -->
<!-- ! usar los crrectas funciones al crear y extraer el token. -->

this is a example git 