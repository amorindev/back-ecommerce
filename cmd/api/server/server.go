package server

import (
	"log"
	"net/http"

	"time"

	v1 "com.fernando/cmd/api/server/v1"
)

// hasta que punto nombrar folderes (no esta muy bien), hasta que punto usar grpcServer, hasta que punto usar server
// y poner alias a las importaciones
type HttpServer struct {
	server *http.Server
}

// debería retornar un error, simpemente log.Fatal
func NewHttpServer(port string) *HttpServer {
	// manejar 	r.Mount("/api/v1", v1.New()), de momento forma manual
	apiV1 := v1.New()

	// limitar request - tiempo de expiracion según el tipo de token

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      apiV1,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := HttpServer{server: serv}
	return &server
}

// close server resources
func (serv *HttpServer) Close() error {
	// TODO: add resource closure
	// me parece que se usa con defer
	//cerrar la base de datos u otros recursos

	// Clear de database postgresql y mongo para empezar nuevamente edebe usar gorutinas
	// solo para desarrrollo
	// o donde llamarlo cuadno cerramos el server?
	//pgConn.Close()

	// !
	/* sessionColl := mongoDB.Collection(config.CollSessions)
	sessionColl.DeleteMany(context.Background(), bson.M{}) */

	return nil
}

func (serv *HttpServer) Start() {
	log.Printf("Http server running http://localhost%s\n", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}

/*
Si quieres eliminar todos los documentos de una colección en MongoDB sin aplicar un filtro específico, simplemente usa un filtro vacío (bson.M{}), así:

filtro := bson.M{} // Filtro vacío elimina todos los documentos

resultado, err := collection.DeleteMany(ctx, filtro)
if err != nil {
    log.Fatal(err)
}

//fmt.Printf("Documentos eliminados: %d\n", resultado.DeletedCount)


*/

/*
Este código eliminará todos los documentos de la colección. Si en su lugar quieres eliminar la colección completa, puedes usar:
err := collection.Drop(ctx)
if err != nil {
    log.Fatal(err)
}

//fmt.Println("Colección eliminada exitosamente")
⚠ Precaución: Asegúrate de que realmente quieres eliminar todos los datos, ya que esta operación no se puede deshacer.


*/
