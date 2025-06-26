package mongo

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	data *Data
	once sync.Once
)

type Data struct {
	DB *mongo.Client
}

func New() *Data {
	once.Do(initDB)
	return data
}

func initDB() {
	db, err := getConnection()
	if err != nil {
		log.Fatal(err)
	}
	data = &Data{
		DB: db,
	}
}

// como manejar los contextos
// deberíamos llamar ping desde main nop
// * principo de demeter no hablar con desconocidos,
// verifaicr funciones deberían ser punteros porque supuestamente no se esta afectando,
// el close podría ser
// ? que context usar en ambos casos
func (data *Data) Ping() error {
	return data.DB.Ping(context.Background(), nil)
}

// ? Como usar esta función?
func (data *Data) Close() error {
	// close que recursis adicionales puedo cerrar ?
	return data.DB.Disconnect(context.Background())
}
// HttpServer ya no tiene dependencia
// con esto me aseguro de no tener varias conecciones

func (data *Data) RemoveDB(onlyData bool){
	
}