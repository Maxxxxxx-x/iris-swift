#!make

APP_EXECUTABLE=iris-swift
APP_PATH=./cmd/api

all: build test

tidy:
	@go mod tidy
	@go vet ./...
	@go fmt ./...

clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ./tmp
	@rm -f coverage*.out


test:
	@echo "Starting test..."
	@go test ./... -v


coverage: test
	@go tool cover -html=coverage.out


build:
	@echo "Building application..."
	@GOARCH=amd64 GOOS=linux go build -o ./tmp/${APP_EXECUTABLE} ${APP_PATH}
	@echo "Build passed"


run: build
	@chmod u+x ./tmp/${APP_EXECUTABLE}
	./tmp/${APP_EXECUTABLE}


watch:
	@if command -v air > /dev/null; then \
		air; \
		echo "Starting air..."; \
	else \
		read -p "Air is not insalled. Install? [Y/n]" choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/air-verse/air@latest; \
			air; \
			echo "Starting air..."; \
		else \
			echo "Air not installed. Exiting..."; \
			exit 1; \
		fi; \
	fi; \
