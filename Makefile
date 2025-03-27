PROJECT := service

BUILD_DIR := ./build
FILES_DIR := ./files

PROJECT_BINARY := $(FILES_DIR)/$(PROJECT)
PROJECT_SRC_PATH := ./cmd/$(PROJECT)

COVER_OUTPUT := $(BUILD_DIR)/coverage.out
COVER_HTML := coverage.html

DOCKER_IMAGE_NAME := app-dev
DOCKERFILE_DIR := ./deployments
GOFLAGS_ARG := -buildvcs=false

.PHONY: build 
.PHONY: clean

build: build_dir lint compile test

lint: 
	GOFLAGS=$(GOFLAGS_ARG) golangci-lint run --timeout 5m -v --config .golangci.yml

test: 
	go test ./... -v -count=1 -race -timeout=1m -covermode=atomic -coverprofile=$(COVER_OUTPUT)
	go tool cover -o $(COVER_HTML) -html=$(COVER_OUTPUT)

compile: build_dir
	GOFLAGS=$(GOFLAGS_ARG) go build -o $(PROJECT_BINARY) $(PROJECT_SRC_PATH)

run:
	$(PROJECT_BINARY)

build_dir: 
	mkdir -p $(BUILD_DIR)
	mkdir -p $(FILES_DIR)

clean: 
	rm -rf $(BUILD_DIR)
	rm -rf $(FILES_DIR)

####DOCKER####

docker-build:
	docker build -f $(DOCKERFILE_DIR)/Dockerfile -t $(DOCKER_IMAGE_NAME) .

docker/%: docker-build
	docker run --rm -it -v $(PWD):/$(PROJECT) -w /$(PROJECT) $(DOCKER_IMAGE_NAME) make $*
