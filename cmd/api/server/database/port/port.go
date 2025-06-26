package port

type Database interface {
	// debe retornar los servicio que se necesita
	GetServices()
}