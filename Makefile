
# Reading env file for environment variables
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# MySQL Connection string generated via environment variables from .env file.
mysql_connection = "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):${MYSQL_PORT})/$(MYSQL_DATABASE)"


default: help

.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done


.PHONY: docker-compose-build
docker-compose-build: # Create docker-compose build
	docker-compose -f docker-compose.yml build $(c)

.PHONY: docker-compose-up
docker-compose-up: # Starts docker-compose first time
	docker-compose -f docker-compose.yml up -d $(c)

.PHONY: docker-compose-start
docker-compose-start: # Starts docker-compose
	docker-compose -f docker-compose.yml start $(c)

.PHONY: docker-compose-down
docker-compose-down: # Stops docker-compose
	docker-compose -f docker-compose.yml down $(c)

.PHONY: docker-compose-destroy
docker-compose-destroy: # Destroys docker-compose
	docker-compose -f docker-compose.yml down -v $(c)

.PHONY: docker-compose-stop
docker-compose-stop: # Temporary stops docker-compose
	docker-compose -f docker-compose.yml stop $(c)

.PHONY: migration-up
up:  # Makes up the migrations for MYSQL
	migrate -path db/migrations -database $(mysql_connection) -verbose up

.PHONY: migration-down
down:  # Makes down the migrations for MYSQL
	migrate -path db/migrations -verbose -database $(mysql_connection) 
