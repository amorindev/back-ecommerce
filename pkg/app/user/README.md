# # Responsabilidades de cada capa

Cada usuario puede tener solo una cuenta por email, ya sea con google apple u otro,

# # # Sign In

<!-- * Handler -->

- auth: password
- user: email
- validar:

<!-- * Service -->

- auth: provider, passwordHash
- user: email_verified

<!-- * * Database -->

- user: id
- auth: id, userid,

<!-- * Service -->

- password = "", junto con omitempty
- passwordHash = "", junto con omitempty

# # # Sign up

<!-- * Hanlder -->

-
-
- validar confirm password

<!-- * todos los servicios auth.password = "" limpiar para todos-->
<!-- * aunque tenga json "-" es importante eliminarlo desde el servicio-->

<!-- ! al crear solo el auth donde se asigna el id service o repo ycómo validarlo -->
<!-- ! es necesario por ejemplo user que esta dentro de auth o auth verificar que no sea nulo antes  -->
<!-- ! de ingresar a sus campos -->

//TODO:validar confirm desde aqui o desde handler? y tambien desde el frontend para no hacer peticiones
//TODO: innecesarias

1. get get by email se diferencian en informacion para sign in y up y el otro datos que pueden ser vistos
2. ver los json y bson etiquetas, dejar solo los servicios necesarios, que combinen los tags junto a hacer nulos los campos
3. ver que se valida en cada capa, y el logguer
4. validaciones de go-l
5. ver que se verifica en cada capa y os punteros del user dentro de auth verificar que no sea nulo antes
   de ingresar a un campo
6. verificar en el handler que vengan los datos del usuario crear sign in y sign up request (ver)

Beneficios
es una capa addicional que nos permite no añadir complegidad al servicio
una colleccion por repositorip
https://www.youtube.com/watch?v=y4Ot7CEpkb8

Auth transactiosn tendrá

// ? que pasa si agregamos uno de mas dará error en ejecución
// si probar por aparte

roleIDs := make([]string, len(roles)-1)
for _, role := range roles {
roleIDs = append(roleIDs, role.ID.(string))
}

<!-- ! falta crear usuarios admin que tendran permission en transaction -->
<!-- ! ver si es necesario al role user asignarle product-read permission -->
<!-- ! si es asi ya no crearemos otra tranasccion sino servirá para ambos -->

validar el email 




que datos debe tener el User al iniciar sesion como email, isverified

get current user