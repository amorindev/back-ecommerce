package model

// TODO: ver el video del profe de como lo pasa relacional a no relacional
// ?pasar a core folder ?
type RoleCreateParams struct {
	Name string `json:"name"`
}

type Role struct {
	// me da error si es estring por que usa decode
	ID   interface{} `json:"id" bson:"_id"`
	Name string      `json:"name" bson:"name"`
}

type UserRole struct {
	ID     interface{} `json:"-" bson:"_id"`           // not null cascade como hacer con transacciones
	UserID interface{} `json:"user_id" bson:"user_id"` // not null cascade como hacer con transacciones
	RoleID interface{} `json:"role_id" bson:"role_id"` // primary
	// que pasa si tengo mas datos aqui
	// * Creo un servicio user_role ?
}
