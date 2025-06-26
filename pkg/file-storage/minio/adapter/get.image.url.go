package adapter

import (
	"context"
	"net/url"
	"time"
)

// TODO fileName cambiar a path o ago porque estamos incluyendo la ruta completa
// TODO o se va a cambiar la estrucutra filestorage dentro de cada Dominio ver de momento fileName
// TODO con toda la ruta

func (a *Adapter) GetImageUrl(ctx context.Context, fileName string) (*url.URL, error) {
	// * Set request parameters for content-disposition.
	reqParams := make(url.Values)
	//reqParams.Set("response-content-type", "image/svg+xml")
	//reqParams.Set("response-content-disposition", "inline")

	// * Generates a presigned url which expires in a day.
	time := time.Hour * 24 * 7 // esyo en prueba time.Hour * 24 //time.Second*24*60*60
	presignedURL, err := a.MinioClient.PresignedGetObject(ctx, "auth-tmpl", fileName, time, reqParams)
	if err != nil {
		return nil, err
	}

	return presignedURL, nil
}
