.PHONY: .run
run:
	go run cmd/url-shortener/main.go

.PHONY: .migrate
migrate:
	go run cmd/migration/main.go

.PHONY: .swag
swag:
	swag init -d cmd/url-shortener,internal/service