package model

import (
	"time"
	//"com.fernando/pkg/app/ecomm/variation-option/model"
)

// ! todos puntero?
// ImgUrl         string   usaremos url no string
type ProductItem struct {
	ID          interface{} `json:"id" bson:"_id"`
	Stock       int         `json:"stock" bson:"stock"`
	Discount    int         `json:"discount" bson:"discount"`
	Rating      int         `json:"rating" bson:"rating"`
	Price       float64     `json:"price" bson:"price"`
	FileName    string      `json:"-" bson:"file_name"`              // uso interno para e backet-  se retorna?
	File        []byte      `json:"-" bson:"-"`                      // se retornna? se debería omitempty o -  json y usar imgurl o fi
	ContentType string      `json:"content_type,omitempty" bson:"-"` // para el init no para la respuesta api
	// !falta img_url o queda el *url.Url
	// * esta bien url.URl es de la biblioteca estandar
	// *url.URL parece que se va
	//FileUrl        *url.URL    `json:"-" bson:"-"` // cual es la diferencia con string parece que se va tener
	ImgUrl    string      `json:"img_url" bson:"-"`
	CreatedAt *time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at" bson:"updated_at"`
	ProductID interface{} `json:"-" bson:"product_id"`
	// * (bson"-") para productos on varaciones
	Options []*Option `json:"options" bson:"options"`
}
type Option struct {
	Name string `json:"name"` // * este es COLOR y SIZE y.....
	// * para insertar en product config necesitamos el id del variation_optin id del Value
	VarOptionID interface{} `json:"-" bson:"-"` // * de momento es como un auxiliar
	Value       string      `json:"value"`      // * ese es RED, YELLOW, M, L
}
