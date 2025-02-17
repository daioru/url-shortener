.PHONY: .run
run:
	go run cmd/url-shortener/main.go

.PHONY: .migrate
migrate:
	go run cmd/migration/main.go

.PHONY: .swag
swag:
	swag init -d cmd/url-shortener,internal/service

.PHONY: .build
build:
	CGO_ENABLED=0  go build \
		-tags='no_mysql no_sqlite3' \
		-o ./bin/url-shortener$(shell go env GOEXE) ./cmd/url-shortener/main.go