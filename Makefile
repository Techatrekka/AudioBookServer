PROJECT_NAME = audiobook-server
DOCKER_COMPOSE = docker-compose

.PHONY: build run clean

build:
	docker build -t $(PROJECT_NAME) .

run:
	$(DOCKER_COMPOSE) up --build

clean:
	$(DOCKER_COMPOSE) down -v

