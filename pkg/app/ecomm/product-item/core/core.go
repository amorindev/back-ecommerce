package core

// la imagen se obtendrá desde headers
type CreateProductReq struct {
	Name             string                `json:"name"`
	Description      string                `json:"description"`
	CategoryID       string                `json:"category_id"`
	Status           string                `json:"status"`
	CreateProductReq *CreateProductItemReq `json:"product"` // seria para productos sin variantes
}

// mejor lo voy a separar en dos estructuras, es
type CreateProductItemReq struct {
	Stock    int     `json:"stock"`
	Discount int     `json:"discount"` // 10%
	Rating   string  `json:"rating"`
	Price    float64 `json:"price"`
}

// TODO se necesita los ids de las variation-options desde la ui
type CreateProductVariantsReq struct {
	// imagen general desde el servicio
	Name                    string                   `json:"name"`
	Description             string                   `json:"description"`
	CategoryID              string                   `json:"category_id"`
	Status                  string                   `json:"status"`
	CreateProductVariantReq *CreateProductVariantReq `json:"variants"`
}

// obtener la imagen del header
type CreateProductVariantReq struct {
	Stock     int64    `json:"stock"`
	Discount  float32  `json:"discount"`
	Rating    string   `json:"rating"`
	Price     float64  `json:"price"`
	OptionIDS []string `json:"options_ids"` // ejem rojo y xl su ids
}

// * todo lo necesario para crear un producto
/* type CreateProductReq2 struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	FileName    string  //`json:"file_name"` // or ImageUrl
	File        []byte
	CategoryID  interface{} `json:"category_id"`
}
*/
