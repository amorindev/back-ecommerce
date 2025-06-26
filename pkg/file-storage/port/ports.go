package port

import (
	"context"
	"net/url"
)

// ! ver el tema de los nombres
type FileStorageAdapter interface {
	// TODO incluir estructura completa products/someting de momento, parece que sirve para 
	// cualquier archivo
	UploadImage(ctx context.Context, fileName string, file []byte, contentType string) error
	GetImageUrl(ctx context.Context, fileName string) (*url.URL, error)
}

// UploadProduct, no me gusta ponle solo upload o uploadfile
// ! ver el sistema de rutas para las imagenes
// * genera el id unico para las imagenes compete a este servicio
type FileStorageSrv interface {
	// ? no sería el modelo si no los []bytes?
	UploadProduct(ctx context.Context, fileName string, file []byte, contentType string) error
	GetImageUrl(ctx context.Context, fileName string) (*url.URL, error)
}
