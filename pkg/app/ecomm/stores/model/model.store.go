package model

import "com.fernando/pkg/app/ecomm/address/model"

// City *string,con esto ya sería complicarno el flujo por que son otras tablas
// y no se si todos los paises tienen provinci distrito
// usar google maps para crear la direccion del store?
// como estore se va a relacionar con los productos
// Y que pasa si el producto que quiere comprar el usuario esta en otro almacén
// me parece que se relacion con Address
// relacionar productID con store
// address tiene un label usar Name?

type Store struct {
	ID          interface{}    `json:"id" bson:"_id"`
	Name        *string        `json:"name" bson:"name"`
	Descripcion *string        `json:"description" bson:"descripcion"`
	AddressAgt  *model.Address `json:"address" bson:"address"`
	/* 	OpeningTime *time.Time  `json:"opening_time" bson:"opening_time"` era fecha estimada de recojo
	   	ClosingTime *time.Time  `json:"closing_time" bson:"closing_time"` */
	// Phone o telf
	// De momento esto type si es sucursal principal almacen si esta activo
	// de moemnto solo eso tambíen podría ser el número
}

// no se si descuento devbe ir dentro del producto
// al  producto agregar costo de
