package model

import (
	"time"

	"com.fernando/pkg/app/ecomm/product-item/model"
)

// TODO: los archivos de los repos tendrán insert convencion de base de datos
// TODO: por que paara buscar un archivo se hace pesado con ctrl+p
// * ImgUrl o FileUrl -
// * como se asigna ImgUrl cuando es sin variantes y con variantes
// ? como será lalogina para asignar una imagen al product sin o con variaciones y que campos tendrá
// ? file name si o si, y no cargar dos veces el archivo ver la lógica
type Product struct {
	ID          interface{} `json:"id" bson:"_id"`
	Name        string      `json:"name" bson:"name"`
	Description string      `json:"description" bson:"description"`
	FileName    string      `json:"-" bson:"file_name"`
	//FileUrl     *url.URL    `json:"url" bson:"-"` // * será asignado del primero de la lista, demomento no usado
	ImgUrl    string     `json:"img_url" bson:"-"`
	Status    string     `json:"status" bson:"status"`
	Sku       string     `json:"sku" bson:"sku"`
	CreatedAt *time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`
	// * vategory id sirve para crear desde el handler pero no lo quiero retornar
	// * tenemos dos struct para crear entonces lo dejaré en -
	// * recordar la parte de cargar desde el json file entonce somitempty despues lo arreglo
	// * desde donde hacerlo = nil desde el service
	CategoryID   interface{} `json:"category_id,omitempty" bson:"category_id"`
	CategoryName string      `json:"category_name,omitempty" bson:"-"` // ? solo para elinsert json
	// * (bson "-")

	ProductItems []*model.ProductItem `json:"product_items" bson:"product_items,omitempty"` // de momento lo vamos a guardar
	// * (bson"-"") configurar los insert,json
	Variations []*Variation `json:"variations" bson:"variants"` // me parece solo para los gets
}

type Variation struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

// * como no te esta saliendo las variciones puedes jugar con los omitempty vetr tags bson
