package model

// ? quien creo el permission seguimiento
type Permission struct {
	ID    interface{} `json:"id" bson:"_id"`
	Name *string     `json:"name" bson:"name"`
}

type RolePermission struct {
	ID           interface{} `json:"-" bson:"_id"`           // not null cascade como hacer con transacciones
	RoleID       interface{} `json:"-" bson:"role_id"`       // not null cascade como hacer con transacciones
	PermissionID interface{} `json:"-" bson:"permission_id"` // not null cascade como hacer con transacciones
	// que pasa si tengo mas datos aqui - quien asigno el permiso userID
	// * Creo un servicio role_permission ?
}
