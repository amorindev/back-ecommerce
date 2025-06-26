package adapter

import (
	"bytes"
	"context"

	"github.com/minio/minio-go/v7"
)

// ! cambia todo
// * a to adp
// * s to srv
// * h hdl o no se

func (a *Adapter) UploadImage(ctx context.Context, fileName string, file []byte, contentType string) error {
	// * que pasa si se crea en la base de datos y no en el fie sotage
	// * no podemos usar transacciones o si? o algo asi,
	// ! primero se crea la imagen ver el medium

	// si se deside asignar un solo backet por ddd del adapter a.BucketName
	// * ver cual es cual
	//fileName := "de"
	//a.Client.PutObject(context.Background(), "products", "products/"+fileName, file)

	fileReader := bytes.NewReader(file)

	//fmt.Printf("Content Type: %s\n", contentType)

	// ! tipo de formato onboarding es video
	options := minio.PutObjectOptions{
		//ContentType: "image/jpg+svg+png",
		//ContentType:        "image/jpg+svg",
		ContentType: contentType,
	}

	// TODO me parece mejor en el handler de cada entidad se crea los servicio y se pasa
	// TODO si necesita un client se le pasa *mongo *minio u otro entonces lo tendríamos
	// TODO mas separado aunque se duplicaria los servicios y pependencias ver o agregar capa

	// *centralazar los nomres o separarlos en cada entidad los nombres de las colleciones los eagrupe
	// * en config pero creo seria mejor separarlos en su entidades

	// * tambien podría ser un backet o carpeta por cada entidad
	// * el nombre del bucket es general puede ir fuera y los nobres de la carpetas /products según DDD
	//fmt.Printf("Size: %v\n", len(file))
	_, err := a.MinioClient.PutObject(ctx, "auth-tmpl", fileName, fileReader, -1, options)
	if err != nil {
		return err
	}

	return nil
}

/* func (a *Adapter) UploadImage(ctx context.Context, fileName string, file []byte) error {
	fileReader := bytes.NewReader(file)

	// ! tipo de formato onboarding es video
	options := minio.PutObjectOptions{
		//ContentType: "image/jpg+svg+png",
		ContentType:        "image/svg",
		ContentDisposition: "inline",
	}

	_, err := a.MinioClient.PutObject(ctx, "ecomm-test", fileName, fileReader, -1, options)
	if err != nil {
		return err
	}

	return nil
}
*/

/* func (a *Adapter) UploadImage(ctx context.Context, fileName string, file []byte, contentType string) error {
	fileReader := bytes.NewReader(file)

	fmt.Printf("Content Type: %s\n", contentType)

	// ! tipo de formato onboarding es video
	options := minio.PutObjectOptions{
		ContentType: contentType,
	}

	_, err := a.MinioClient.PutObject(ctx, "ecomm-test", fileName, fileReader, -1, options)
	if err != nil {
		return err
	}

	return nil
} */
