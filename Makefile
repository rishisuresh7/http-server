build:
	@echo "Building binary..."
	@go build -ldflags "-X main.Version=1.0.0" -o build/http-server main/main.go
	@echo "Done"

run:
	@echo "Running http-server..."
	@go run main/main.go

clean:
	@echo "Cleaning build folder..."
	@rm -rf build
	@echo "Done"

test:
	@echo "Running tests..."
	@go test ./...
