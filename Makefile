postgres:
	docker run --name ecommerce -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.2

createdb:
	docker exec -it ecommerce createdb --username=root --owner=root ecommerce

dropdb:
	docker exec -it ecommerce dropdb ecommerce

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/ecommerce?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/ecommerce?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  github.com/zahrou/ecommerce/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown test server mock
	