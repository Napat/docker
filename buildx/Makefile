# Docker image name and tag
IMAGE_NAME := my-go-app
VERSION := 1.0.0
REGISTRY := docker.io/napat

.PHONY:all
# Build, ship multi-architecture images amd run container
all: create-builder build-multi compose-up

.PHONY:create-builder
# Create a new builder instance
create-builder:
	docker buildx inspect mybuilder >/dev/null 2>&1 || \
		docker buildx create --name mybuilder --driver docker-container --bootstrap
	docker buildx use mybuilder

.PHONY:build-multi
# Build multi-architecture images
build-multi:
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		-t $(REGISTRY)/$(IMAGE_NAME):$(VERSION) \
		-t $(REGISTRY)/$(IMAGE_NAME):latest \
		--push \
		.

.PHONY:build-local
# Build multi-architecture images locally (without pushing)
build-local:
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		-t $(IMAGE_NAME):latest \
		--load \
		.

.PHONY:compose-up
# Run Docker Compose
compose-up:
	docker-compose up -d

.PHONY:clean
# Clean up builder
clean:
	docker buildx rm mybuilder