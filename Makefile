build:
	@echo "Building binary..."
	@go build -ldflags "-X main.Version=1.0.0" -o build/http-server main/main.go
	@echo "Built at build/http-server"

build-linux:
	@echo "Building binary for linux..."
	@GOOS=linux GOARCH=arm64 go build -ldflags "-X main.Version=1.0.0" -o build/http-server main/main.go
	@echo "Built at build/http-server"

run:
	@echo "Running http-server..."
	@source script.sh && go run main/main.go

clean:
	@echo "Cleaning build folder..."
	@rm -rf build
	@echo "Done"

test:
	@echo "Running tests..."
	@go test ./...
