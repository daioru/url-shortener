.PHONY: .run
run:
	go run cmd/main.go

.PHONY: .migrate
migrate:
	go run internal/migrations/main.go