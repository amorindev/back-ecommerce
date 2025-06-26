Tengo las siguientes coleciones en mongo y golang, modelo para hacer
productos con variaciones

Product
\_id
680836446a2b7f5dd8652f4b
name
"Zapatillas New Athletic Casual Dauscole Mujer"
description
"Comodidad y estilo se unen en este modelo casual, ideal para el día a …"
file_name
"products/black-36.jpg"
status
"active"
created_at
2025-04-23T00:37:23.244+00:00
updated_at
2025-04-23T00:37:23.244+00:00
category_id
6807b7e53ec7725d59a6505e

Cada Product tiene muchos productos

\_id
680836446a2b7f5dd8652f4c
stock
10
discount
0
rating
5
price
100
file_name
"products/black-36.jpg"
created_at
2025-04-23T00:37:23.244+00:00
updated_at
2025-04-23T00:37:23.244+00:00
product_id
680836446a2b7f5dd8652f4b

\_id
680836446a2b7f5dd8652f4d
stock
10
discount
0
rating
5
price
110
file_name
"products/black-37.jpg"
created_at
2025-04-23T00:37:23.244+00:00
updated_at
2025-04-23T00:37:23.244+00:00
product_id
680836446a2b7f5dd8652f4b

tambien tenemos la variaciones como size y color

\_id
6807b7e53ec7725d59a65060
name
"COLOR"

las variaciones tienen opciones en tallas tenemos 41 42 43 color rojo blanco

\_id
6807b7e53ec7725d59a6506a
value
"40"
variation_id
6807b7e53ec7725d59a65061

\_id
6807b7e53ec7725d59a65063
value
"BLACK"
variation_id
6807b7e53ec7725d59a65060

Para relacionar ambos tenemos la tabla o collecion itermedia

\_id
680836446a2b7f5dd8652f52
product_id
680836446a2b7f5dd8652f4c
var_option
6807b7e53ec7725d59a65063

\_id
680836446a2b7f5dd8652f53
product_id
680836446a2b7f5dd8652f4c
var_option
6807b7e53ec7725d59a65066

quiero hacer la consulta para este formato de api

type Product struct {
ID interface{} `json:"id" bson:"_id"`
Name string `json:"name" bson:"name"`
Description string `json:"description" bson:"description"`
FileName string `json:"-" bson:"file_name"`
FileUrl _url.URL `json:"img2_url" bson:"-"` // _ será asignado del primero de la lista
ImgUrl string `json:"img_url" bson:"-"`
Status string `json:"status" bson:"status"`
CreatedAt *time.Time `json:"created_at" bson:"created_at"`
UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`
CategoryID interface{} `json:"category_id" bson:"category_id"`
CategoryName string `json:"category_name,omitempty" bson:"-"`
Products []*model.Product `json:"products" bson:"-"` // de momento lo vamos a guardar
Variations []*Variation `json:"variations" bson:"-"` // me parece solo para los gets
}

type Variation struct {
Name string `json:"name"`
Values []string `json:"values"`
}

"variations": [
{
"name": "COLOR",
"values": [
"BLACK",
"LEAD"
]
},
{
"name": "SIZE",
"values": [
"36",
"37",
"38"
]
}
]

y en el producto
type Product struct {
	ID       interface{} `json:"id" bson:"_id"`
	Stock    int         `json:"stock" bson:"stock"`
	Discount int         `json:"discount" bson:"discount"`
	Rating   int         `json:"rating" bson:"rating"`
	Price    float64     `json:"price" bson:"price"`
	FileName string      `json:"-" bson:"file_name"` // uso interno para e backet-  se retorna?
	File     []byte      `json:"-" bson:"-"`         // se retornna? se debería omitempty o -  json y usar imgurl o fi
	// !falta img_url o queda el *url.Url
	// * esta bien url.URl es de la biblioteca estandar
	FileUrl        *url.URL    `json:"file_url" bson:"-"` // cual es la diferencia con string
	ImgUrl         string      `json:"img_url" bson:"-"`
	CreatedAt      *time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt      *time.Time  `json:"updated_at" bson:"updated_at"`
	ProductID interface{} `json:"-" bson:"product_id"`
	// * para productos on varaciones
	Options []*Option `json:"options" bson:"-"`
}
type Option struct {
	Name string `json:"name"` // * este es COLOR y SIZE y.....
	// * para insertar en product config necesitamos el id del variation_optin id del Value
	VarOptionID interface{} `json:"-" bson:"-"` // * de momento es como un auxiliar
	Value       string      `json:"value"`      // * ese es RED, YELLOW, M, L
}
sin el VarOptionID interface{}