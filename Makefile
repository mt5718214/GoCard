postgres:
	docker run --name postgres13 -e POSTGRES_USER=gocard -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:13

# https://stackoverflow.com/questions/26992821/postgresql-how-to-insert-null-value-to-uuid
addplugin:
	docker exec postgres13 psql -U gocard -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"

createdb:
	docker exec postgres13 createdb --username=gocard --owner=gocard gocard

createtestdb:
	docker exec -it postgres13 createdb --username=gocard --owner=gocard gocard_test && docker exec -it postgres13 psql -U gocard -d gocard_test -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";" && migrate -path db/migration -database "postgres://gocard:secret@localhost:5432/gocard_test?sslmode=disable" --verbose up

dropdb:
	docker exec postgres13 dropdb --username=gocard gocard

migrateup:
	migrate -path db/migration -database "postgres://gocard:secret@localhost:5432/gocard?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgres://gocard:secret@localhost:5432/gocard?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

swag:
	swag init

.PHONY: postgres addplugin createdb dropdb migrateup migratedown sqlc test swag
