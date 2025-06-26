package port

import "context"

// * Adapter para interactuar con servicios internos no lo pongo fuera
// * como el email por que sirve solo para auth

// * los primitivos tienen una funcion validate para que no vengas datos con vaor ceo o nulos
// ? de igual manera sacar una funcion en el hander para el google validate token?
// * o saco un  adapter como este, si interacciona con servicios de google si es mejor adaoper
type AuthAdapter interface {
	GoogleValidateToken(ctx context.Context, token string) (string, error)
	AppleValidateToken(ctx context.Context, token string) (string, error)
	// deberia devolcer user o FacebookUserClaims
	FacebookValidateToken(ctx context.Context, token string) (string, error)
}
