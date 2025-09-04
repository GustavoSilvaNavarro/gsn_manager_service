#!make
PORT = 8080
SERVICE_NAME = gsn_manager_service
CONTAINER_NAME = $(SERVICE_NAME)
DOCKER_COMPOSE_TAG = $(SERVICE_NAME)_1
TICKET_PREFIX := $(shell git branch --show-current | cut -d '_' -f 1)

up:
	go run ./src/main.go
