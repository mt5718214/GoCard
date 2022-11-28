postgres:
	docker run --name postgres13 -e POSTGRES_USER=gocard -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:13

addplugin:
	docker exec -it postgres13 psql -U gocard -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"

createdb:
	docker exec -it postgres13 createdb --username=gocard --owner=gocard gocard

dropdb:
	docker exec -it postgres13 dropdb gocard

migrateup:
	migrate -path db/migration -database "postgres://gocard:secret@localhost:5432/gocard?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgres://gocard:secret@localhost:5432/gocard?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres addplugin createdb dropdb migrateup migratedown sqlc test
