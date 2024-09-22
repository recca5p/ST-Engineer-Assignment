DB_URL=postgresql://myuser:mypassword@localhost:5432/demo?sslmode=disable

infras:
	docker-compose up -d

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination=db/mock/store.go -package=mockdb Golang-Bank-self-study-/db/sqlc Store

.PHONY: infras migrateup migratedown new_migration sqlc test server mock