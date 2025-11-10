.PHONY: server migrate force_migrate fmt

fmt:
	go fmt ./cmd/api/
	
server:
	go run ./cmd/api/main.go

migrate:
	migrate -database postgresql://postgres:postgres@localhost:5432/oplan_dev?sslmode=disable \
	-path ./internal/database/migrations/ \
	-verbose $(where)

force_migrate:
	migrate -database postgresql://postgres:postgres@localhost:5432/oplan_dev?sslmode=disable \
	-path ./internal/database/migrations/ \
	-verbose force 1

sqlc:
	sqlc generate