APP_NAME := clario-weather-cli
BUILD_DIR := build

.PHONY: build run test clean

build:
	@echo "==> Building $(APP_NAME)..."
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/weather-cli

run: build
	@echo "==> Running $(APP_NAME)..."
	./$(BUILD_DIR)/$(APP_NAME)

test:
	@echo "==> Running tests..."
	go test -v ./...

clean:
	@echo "==> Cleaning up..."
	rm -rf $(BUILD_DIR)
