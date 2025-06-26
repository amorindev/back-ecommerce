package model

import "time"

// *  necesitamos que el usuario no lo vuelva a ver
// * 1. cron job que elimine ek onboarding pasado un  tiempo
// * 2. marcar como visto un onboarding desde donde hacerlo frontend o backend?
// * Historia de usuario - los usuarios no volveran a ver los onbording vistos
// * ver el flujo para mostrarlo al profesor a grabar video
// * u onboardign pasados la checha de expiracion
// * me parece mejor seem para ver si el usuario lo avisto
// * ver que campose para el administrador de momento no lo necesito (expires_at tambin)
// * en la pantalla de admin ver el objeto creado puede ser omitempty
// * si es user lo asignamos "" o null si no es admin lo devolvemos
// * Expires debe ser mayor a la fecha en la que se crea service verificar
// ! Seen muchos a muchos para marcar como visto por cada usuaio
// ! ademas al hacer get filtrarlo si fue visto de momento no lo necesitamos
// ! por defecto sera false seen
// ! lo necesito para el init data `json:"file_name" pero usar omitepti para no mostrar en la api
// ! o seria mejor crear otra entidades para no mesclar
// ! si va ser onbaoarding por usuario getall por userid
type Onboarding struct {
	ID          interface{} `json:"id" bson:"_id"`
	FileName    string      `json:"file_name,omitempty" bson:"file_name"` // uso interno para e backet-  se retorna? DB
	File        []byte      `json:"-" bson:"-"`                           // se debería omitempty o -  json y usar imgurl o fi
	ContentType string      `json:"content_type,omitempty" bson:"-"`      // para el init no para la respuesta api
	// me parece mejor FIleurl
	ImgUrl string `json:"img_url" bson:"-"`
	Title  string `json:"title" bson:"title"`
	Text   string `json:"text" bson:"text"`
	//Seen      bool        `json:"-" bson:"seen"` // mejor no retornarlo uso interno es el filtro
	ExpiresAt *time.Time `json:"expires_at" bson:"expires_at"`
	CreatedAt *time.Time `json:"created_at" bson:"created_at"`

	// ver el tema de punteros en structiras
	// no se si poner updated no
	// FileExtension enum para distinguir si es imagen o video,
	// pero tambien se puede obtener del nombre archivo o path
}
