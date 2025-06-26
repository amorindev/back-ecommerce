package service

import (
	"context"
	"errors"
	"time"

	"com.fernando/pkg/app/ecomm/products/model"
)

// TODO despues de esto revisar los ids de los repos de los 3 que se crearon
func (s *Service) Create(ctx context.Context, product *model.Product) error {
	if len(product.ProductItems) == 0 {
		return errors.New("product-items debe tener por lo menos un elemento")
	}

	now := time.Now().UTC()

	product.CreatedAt = &now
	product.UpdatedAt = &now

	bucketFolderStruct := "products/"

	// ? Como asignar la imagen si tiene variantes o no tiene
	// asignando la primera imagen del producto item de la lista al product
	// como se que va tener la misma imagen del producto
	// ya no es necesario guardarlo solo asignarle el nombre
	product.FileName = bucketFolderStruct + product.ProductItems[0].FileName

	product.Sku = "test-sku"

	for _, product := range product.ProductItems {
		product.CreatedAt = &now
		product.UpdatedAt = &now

		//!Como obtengo la url o como se crea para que el usario accedda
		//continue
		//fmt.Println(product.FileName)
		//fmt.Println(bucketFolderStruct)
		product.FileName = bucketFolderStruct + product.FileName
		err := s.FileStorageSrv.UploadProduct(context.Background(), product.FileName, product.File, product.ContentType)
		if err != nil {
			return err
		}
	}

	// * como obtener la url, o quería por que al momento de crear un producto queiro retornar el producto
	// * creado
	//log.Fatal("test")
	/* url := &url.URL{
	}*/

	err := s.ProductTx.Create(context.Background(), product)
	if err != nil {
		return err
	}

	return nil
}

//FileName string      `json:"-" bson:"-"` // uso interno para el nobre del bucket y base de datos
//File     []byte      `json:"-" bson:"-"` // se obtiene desde el header no se retorna ni se guarda DB AUX interno
// !falta img_url o queda el *url.Url
// * esta bien url.URl es de la biblioteca estandar
//FileUrl        *url.URL    // cual es la diferencia con string
// ProductID interface{} `json:"-" bson:"product_id"` // ? Dónde se asigna
// * para productos on varaciones
