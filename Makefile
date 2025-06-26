include .env.dev

.PHONY: Makefile for dev
run:
	@go run main.go

.PHONY: rehacer
compose:
	@docker-compose -f docker-compose.yml up -d

.PHONE: pasar el env file, esta bien por defecto dockercompose lee el .env
.PHONY: @docker-compose -f docker-compose.dev.yml up -d
compose-dev:
	@docker-compose -f docker-compose.dev.yml --env-file .env.dev up

.PHONY: que use todos los protos tambien con utils y timestamp como hacerlo
.PHONY: ? es necesario agregar el timestamp como flutter lo requiere
gen:
	@protoc \
	--proto_path=files/protobuf/v1/auth \
	--go_out=pkg/services/app/auth/grpc/ gen --go_opt=paths=source_relative \
	--go-grpc_out=pkg/services/app/auth/grpc/gen \
	--go-grpc_opt=paths=source_relative files/protobuf/v1/auth/*.proto

ngrok-http:
	@ngrok http http://localhost:$(HTTP_SERVER_PORT)

cloudflared-minio:
	@cloudflared tunnel --url http://localhost:$(MINIO_API_PORT) 2>&1 | grep -o 'https://[^ ]*'

