# Variables
MIGRATIONS_DIR := migrations
GOOS := linux
VERSION ?= latest  # Default version if none is provided
CONTAINER_NAME := stockzilla-service-stock-trader
IMAGE_NAME := vickeyshrestha/service-stock-trader

# Targets
.PHONY: build docker-build

# Clean and build the Go binary
build:
	go clean
	set GOOS=$(GOOS)
	go build -a -installsuffix cgo -o service-stocktrader cmd/main.go

# Build the Docker image with the provided version
# usage eg: make docker-build VERSION=1.0.0
docker-build: build
	docker build -t $(IMAGE_NAME):$(VERSION) .
    @echo "Docker image $(IMAGE_NAME):$(VERSION) built successfully."

# Run the Docker container with the built image
docker-run:
	docker run --rm --name $(CONTAINER_NAME) \
		-p 5112:8085 \
		-p 5113:8086 \
		$(IMAGE_NAME):$(VERSION)
	@echo "Docker container $(CONTAINER_NAME) running with ports 5112:8085 and 5113:8086."

# Push the Docker image to the registry
push: docker-build
	docker push $(IMAGE_NAME):$(VERSION)
	@echo "Docker image $(IMAGE_NAME):$(VERSION) pushed successfully."

#goose create migration
goose-create-migration:
	goose -s -dir $(MIGRATIONS_DIR) create stock_trader sql

goose-up:
	goose -dir $(MIGRATIONS_DIR) postgres "postgres://postgres:admin@localhost:5104/stockzilla?sslmode=disable" up


goose-down:
	goose -dir $(MIGRATIONS_DIR) postgres "postgres://postgres:admin@localhost:5104/stockzilla?sslmode=disable" down
