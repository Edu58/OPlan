.PHONY: server migrate force_migrate fmt

fmt:
	go fmt -n ./...

test:
	go test -v -cover ./...

test-unit:
	go test -short -v -cover ./...

server:
	go run ./cmd/web/main.go

race:
	go test -race ./...

# copylock/CopyLock check
vet:
	go vet ./...

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
