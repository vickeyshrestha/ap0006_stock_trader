# Variables
DOCKER_COMPOSE := docker-compose -f docker-compose.yaml

# Targets
.PHONY: up down

# Start the Docker Compose services
docker-up:
	$(DOCKER_COMPOSE) up -d

# Stop the Docker Compose services
docker-down:
	$(DOCKER_COMPOSE) down
