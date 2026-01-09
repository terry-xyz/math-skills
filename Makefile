BINARY_NAME := math-skills
BINARY_PATH := bin/$(BINARY_NAME)
MAIN_PKG := ./cmd/math-skills

.PHONY: help build run test clean

help:
	@echo "Available targets:"
	@echo "  make build  - Compile binary to bin/"
	@echo "  make run    - Run with data.txt"
	@echo "  make test   - Run all tests"
	@echo "  make clean  - Remove build artifacts"

build:
	go build -o $(BINARY_PATH) $(MAIN_PKG)

run: build
	$(BINARY_PATH) data.txt

test:
	go test -v $(MAIN_PKG)

clean:
	rm -rf bin/
