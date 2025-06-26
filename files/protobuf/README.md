<!-- ! Files sera ignorado por .gitignore asi que no poner .html que se usa para el envío de emails -->

los protos deberian ir separados ya que no son archivos golang y son generados, igua a los sql
y en que afectaría si dejo en services en su propio servico

<!--* Ahora estamos en un estado dediseñar entones comoestamos agregando user dentro de auth -->
<!--* y no estamos creando un user por aprte para auth entonces lo mismo haremos en grpc -->
<!--* no se si sigue DDD, y si es molular veremos más adelante -->
<!--* user en  -->

```go
type Auth struct {
    ID interface{} `json:"-" bson:"_id,omitempty"` // omitempty, "-"
    UserID interface{} `json:"-" bson:"user_id"` // !flujo para el userID
    Provider *string `json:"provider" bson:"provider"` // es visible?
    UserAgregate *userModel.User `json:"user" bson:"user,omitempty"` // se debe validar que no sea nil?
    CreatedAt *time.Time `json:"-" bson:"created_at"`
    UpdatedAt *time.Time `json:"-" bson:"updated_at"`
}
```
// midlewares para todos , en especial para authchanges o current user cual es la diferencia
