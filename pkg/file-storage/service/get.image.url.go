package service

import (
	"context"
	"net/url"
)

// * que seria mejor traer todas las imagenes y asignarlas y si fata consultar y si no interna server
// * error, de momento solo hare varias consultas con un for

func (s *Service) GetImageUrl(ctx context.Context, fileName string) (*url.URL, error) {
	presignedURL, err :=  s.FileStgAdp.GetImageUrl(ctx, fileName)
	if err != nil {
	//? para el retorno (string, error) debe ser *string
	  return nil, err
	}
	//return presignedURL.String(), err
	return presignedURL, err
}
