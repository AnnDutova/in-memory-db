PROJECT := service

BUILD_DIR := ./build
FILES_DIR := ./files


PROJECT_BINARY := $(FILES_DIR)/$(PROJECT)
PROJECT_SRC_PATH := ./cmd/$(PROJECT)

COVER_OUTPUT := $(BUILD_DIR)/coverage.out

.PHONY: build 
.PHONY: clean

build: build_dir lint compile test run

lint: 
	golangci-lint run --timeout 5m -v --config .golangci.yml

test: 
	go test ./... -v -count=1 -race -timeout=1m -covermode=atomic -coverprofile=$(COVER_OUTPUT)

compile: build_dir
	go build -o $(PROJECT_BINARY) $(PROJECT_SRC_PATH)

run:
	$(PROJECT_BINARY)

build_dir: 
	mkdir -p $(BUILD_DIR)
	mkdir -p $(FILES_DIR)

clean: 
	rm -rf $(BUILD_DIR)
	rm -rf $(FILES_DIR)