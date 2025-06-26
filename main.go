package main

import (
	"log"
	"os"

	httpSrv "com.fernando/cmd/api/server"
	//grpcSrv "com.fernando/cmd/grpc/server"
	//"github.com/joho/godotenv"
)

func main() {

	// de momento solo lo comento
	/* err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	} */

	//defer postgresql.CloseConnection(data.DB) esto desde v1 me parece correcto

	hsp := os.Getenv("HTTP_SERVER_PORT")
	if hsp == "" {
		log.Fatal("environment variable HTTP_SERVER_PORT is not set")
	}

	httpServer := httpSrv.NewHttpServer(hsp)
	httpServer.Start()

	// * Si debería dejar usar los dos servidores solo se expone el puerto que usa ngrok

	/* go httpServer.Start()

	gsp := os.Getenv("GRPC_SERVER_PORT")
	if gsp == "" {
		log.Fatal("environment variable GRPC_SERVER_PORT is not set")
	}
	gRPCServer := grpcSrv.NewgRPCServer(gsp)
	gRPCServer.Start() */

}
