package adapter

import (
	"com.fernando/pkg/file-storage/port"
	"github.com/minio/minio-go/v7"
)

var _ port.FileStorageAdapter = &Adapter{}

type Adapter struct {
	// asignar una carpeta igual a mongo collection?
	MinioClient *minio.Client
}

func NewAdapter(client *minio.Client) *Adapter {
	return &Adapter{
		MinioClient: client,
	}
}
