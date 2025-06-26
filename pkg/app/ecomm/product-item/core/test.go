package core

import "time"

// * en este escenario ProductItemAgregate *ProductItem `json:"product_items"`
// * ya deberíamos tener e diseño de la base de datos si deseamos no poner anidados
// * simplemente usamos bson"-" y en los repós ya no haremos auth.useragregate = nil e insertar
type Product struct {
	ID          interface{} `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	// en el backend lo asignamos que no sea requerdio le asignaremos el primero de la lista
	UrlImg              string         `json:"url_img"` // imagen de presentacion ui es para el response
	CreatedAt           *time.Time     `json:"created_at"`
	UpdatedAt           *time.Time     `json:"updated_at"`
	ProductItemAgregate []*ProductItem `json:"product_items"`
	Variations          []Variation    `json:"variations"`
	// ? categoría? o sola para crear
}

type ProductItem struct {
	ID        interface{} `json:"id"`
	Price     float64     `json:"price"`
	Stock     int64       `json:"stock"`
	Discount  int         `json:"discount"` // 10%
	Rating    int         `json:"rating"`
	ImgUrl    string      `json:"img_url"`
	CreatedAt time.Time   `json:"created_at"`
	ProductID interface{} `json:"product_id"`
	Options   []Option    `json:"options"`
}

type Option struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Variation struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
