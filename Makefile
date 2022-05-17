include .env
export

openapi: openapi_http

openapi_http:
	@./scripts/openapi-http.sh http_api internal/ports/http_api http_api

build: go-mod-tidy build-backend build-cli

go-mod-tidy:
	go mod tidy

build-backend:
	go build -ldflags "-s -w" -o cmd/backend main.go

build-cli:
	go build -ldflags "-s -w" -o cmd/client internal/cli/main.go

install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

apply-migrations:
	migrate -path migrations -database "$(BJ_DSN)?x-migrations-table=jack_black_migrations&sslmode=disable" up
