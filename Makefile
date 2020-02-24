.PHONY: migrate
migrate:
	./goose up

.PHONY: drop
drop:
	./goose down

.PHONY: runserver
runserver:
	go run cmd/goapi/main.go
