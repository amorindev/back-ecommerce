package service

import (
	"com.fernando/pkg/file-storage/port"
)

var _ port.FileStorageSrv = &Service{}

// ! adapter es adp no otro revisar

type Service struct {
	FileStgAdp port.FileStorageAdapter
}

func NewFileStgSrv(fileStgAdp port.FileStorageAdapter) *Service{
	return &Service{
		FileStgAdp: fileStgAdp,
	}
}