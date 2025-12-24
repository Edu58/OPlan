.PHONY: server migrate templ tailwind dev build force_migrate fmt

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

templ:
	@echo "Generating Templ templates..."
	@templ fmt ./internal/frontend
	@templ generate

tailwind:
	@echo "Building Tailwind CSS..."
	@npx tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css

tailwind-watch:
	@echo "Watching Tailwind CSS..."
	@npx tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css --watch

# Build for production
build: templ
		@echo "Building for production..."
		@npx tailwindcss -i ./web/static/css/input.css -o ./web/static/css/output.css --minify
		@go build -ldflags="-s -w" -o bin/oplan cmd/web/main.go
		@echo "Build complete! Binary: bin/oplan"

# Run development server with hot reload
dev: templ tailwind
	@echo "Starting development server..."
	@air

# Run without hot reload
run: templ tailwind
	@echo "Running application..."
	@go run cmd/web/main.go

# Production start
prod: build
	@echo "Starting production server..."
	@./bin/oplan
