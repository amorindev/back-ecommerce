package service

import (
	"context"
)

func (s *Service) UploadProduct(ctx context.Context, fileName string, file []byte, contentType string) error {

	// TODO crear el id unico para las imágenes o desde el handler
	return s.FileStgAdp.UploadImage(ctx, fileName, file, contentType)
}
