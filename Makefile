postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root todo_db

dropdb:
	docker exec -it postgres12 dropdb todo_db

create_init_migration:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/todo_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/todo_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/emmyvera/go_todo/db/sqlc Store

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown create_init_migration server