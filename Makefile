.PHONY: migrate
migrate:
	docker run --network host migrator:latest -path=/migrations/ -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

.PHONY: drop
drop:
	docker run --network host migrator:latest -path=/migrations/ -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" drop

.PHONY: runserver
runserver:
	go run cmd/goapi/main.go
