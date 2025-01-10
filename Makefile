DB_URL=postgresql://root:root@localhost:5433/food?sslmode=disable

postgres:
	docker run --name postgres16 --network food-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root food

dropdb:
	docker exec -it postgres16 dropdb food

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate
	
test:
	go test -v -cover -short ./...

server:
	go run main.go

proto:
	rm -f pb/*.pb.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=food-backend \
	--experimental_allow_proto3_optional \
	proto/*.proto
	rm -f doc/statik/statik.go
	statik -src=./doc/swagger -dest=./doc

.PHONY: proto