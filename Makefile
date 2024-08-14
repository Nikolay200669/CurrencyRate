BINARY_NAME=CurrencyRate
BUILD_DIR=build

.PHONY: all build clean run

all: build

build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/currency-app/main.go

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

run:
	@echo "Running..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

test:
	@echo "Running tests..."
	@go test ./...

lint:
	@echo "Linting..."
	@golangci-lint run

enable-autostart:
	@echo "Enabling autostart..."
	@./$(BUILD_DIR)/$(BINARY_NAME) --enable-autostart

disable-autostart:
	@echo "Disabling autostart..."
	@./$(BUILD_DIR)/$(BINARY_NAME) --disable-autostart
