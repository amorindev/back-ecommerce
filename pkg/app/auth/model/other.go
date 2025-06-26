package model

// si usamos transaccioens
// ademas se repitará el email, usar omitempty para no repetir email pero se puede sacar de auth.User.Email
// asi como se creará campo de likes par no estar calulando cada ves
// 1 todo esto es en pro de la veocidad
// otra cosa es que el Auth y el user se crean juntos
// en cambio User y Post de red social se crean por separado anaizar
// ver agregacion en mongo docs parecido a Exec
// se puede cambiar de email? una cuenta y como mantener la consistencia mas adeannteeee!
// evitar joins y como manejar datos repetidos?
// entonces teneoms dos enfoque hacerlo solo en create o separar create account y user
// go senior parece que o une
// ver de nuevo video de transacciones, en una implementacion  maspura no se usa dos colleciones en el
// repository
// si es embedding no hay necesisdad de agregar dos collectioss, per sigue tenniendo ,mas de una responsabilidad
// ver siva separado o ambos  - puede ser ambos usando mongo embeddings y metadata en postgresql
// una colleccion por cada repositio decision purista
type Auth2 struct {
	ID interface{}
}

type User2 struct {
	ID interface{}
	AuthID interface{} // ? o estring o primite?
}
