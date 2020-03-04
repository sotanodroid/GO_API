.PHONY: migrate
migrate:
	./goose up

.PHONY: drop
drop:
	./goose down

.PHONY: runserver
runserver:
	go run cmd/goapi/main.go

.PHONY: test
test:
	go test ./... -v -covermode=count

.PHONY: lint
lint:
	golangci-lint run -D errcheck
