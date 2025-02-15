include .env
export $(shell sed 's/=.*//' .env)

SHELL := /bin/bash
GO_ENV_PATH = $(shell go env GOPATH)

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)


setup:
	@echo ${YELLOW}Installing Golangci-lint${WHITE}
	sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GO_ENV_PATH)/bin v1.54.2
	@echo ${GREEN}Golangci-lint installed successfully
	@echo ''

	@echo ${YELLOW}Installing atlasgo${WHITE}
	curl -sSf https://atlasgo.sh | sh
	@echo ${GREEN}Atlasgo installed successfully
	@echo ''

	@echo ${YELLOW}Installing dependencies${WHITE}
	go mod download
	@echo ${GREEN}Dependencies installed successfully
	@echo ''

	@echo ${YELLOW}Installing gotestfmt${WHITE}
	go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@latest
	@echo ${GREEN}gotestfmt installed successfully
	@echo ''

	@echo ${YELLOW}Installing Mockery${WHITE}
	go install github.com/vektra/mockery/v2@v2.52.1
	@echo ${GREEN}Mockery installed successfully
	@echo ''

all: build test

build:
	@echo "Building..."
	
	
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

	
# Create DB container
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v
# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

.PHONY: all build run test clean watch docker-run docker-down itest

migrations-status:
	@atlas migrate status --dir file://migrations --url $(DSN)

migrations-generate:
	@atlas migrate hash
	@atlas migrate diff --env gorm
	@echo ${GREEN}Migration created successfully

migrations-execute:
	@atlas migrate apply --env gorm -u $(DSN)
