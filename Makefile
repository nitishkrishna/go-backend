postgres:
	docker run --name postgres-bookstore -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mypgpassword -d postgres:12-alpine

createdb:
	docker exec -it postgres-bookstore createdb --username=root --owner=root bookstore 

dropdb:
	docker exec -it postgres-bookstore dropdb bookstore

migrateup:
	migrate -path db/migration -database "postgresql://root:mypgpassword@localhost:5432/bookstore?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mypgpassword@localhost:5432/bookstore?sslmode=disable" -verbose down

createdb-catalog:
	docker exec -it postgres-bookstore createdb --username=root --owner=root catalog

dropdb-catalog:
	docker exec -it postgres-bookstore dropdb catalog

build:
	go build -mod=vendor -a -o app -ldflags "-X main.version=latest" cmd/main.go

run-backend:
	go run cmd/main.go

run-ui:
	cd ui && yarn && yarn dev


.PHONY: postgres createdb dropdb migrateup migratedown createdb-catalog dropdb-catalog build run-backend run-ui
