package minio

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// ! si queda minio por sepaado podira ir dentro de file-storage/minio
type MinioClient struct {
	Client *minio.Client
}

// retornar un struct como microblog o el cliente como ahora

// * estacreando dos veces el client para init y para el v1 como el init se usa solo una ves
// * no habrá problema  solo es un aviso
func NewClient() *MinioClient {

	// verificar si es nulo
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretKeyID := os.Getenv("MINIO_SECRET_KEY")

	// ? ponerlo como .env -si para no cambiarrlo en produccuib
	useSSL := true // * en railway podría ser true  por que es https o ngrok

	// * Initialize minio client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretKeyID, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatal("Minio client failed: %w", err)
	}

	newMinioClient := &MinioClient{
		Client: minioClient,
	}

	return newMinioClient
}

// * I set the Secure as false because i don’t use HTTPS. You can adjust it to true if you use HTTPS.
// ! En produccion no seria true  por que usas railway
// ! me parece que no data error si usas false pero será inseguro

// no se si usar pointer *MinioClient
/* func (minioClient *MinioClient) CreateBuckets() {

	minioClient.Client.MakeBucket()
} */

func (client *MinioClient) CreateStorage() {
	// crear bucket desde aqui o como base de datos desde el config
	bucketName := os.Getenv("MINIO_BUCKET_NAME")
	if bucketName == "" {
		log.Fatal("environment variable MINIO_BUCKET_NAME is not set")
	}

	found, err := client.Client.BucketExists(context.Background(), bucketName)
	if err != nil {
		//log.Fatalf("create Storage Bucket exists err: %v\n", err)
	}
	if found {
		fmt.Printf("Bucket exists.\n")
	} else {
		//fmt.Printf("Bucket does not exists.\n")
		err := client.Client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			fmt.Printf("create Storage failed: %s\n", err.Error())
		}
	}
}
