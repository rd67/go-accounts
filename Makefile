
# Reading env file for environment variables
ifneq (,$(wildcard ./app.env))
    include app.env
    export
endif

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

# .SILENT:
.PHONY: migration-up
migration-up:  # Makes up the migrations for DB
	migrate -path db/migrations -database "$(DB_SOURCE)" -verbose up

# .SILENT:
.PHONY: migration-down
migration-down:  # Makes down the migrations for DB
	migrate -path db/migrations -database "$(DB_SOURCE)" -verbose down

.PHONY: sqlc
sqlc:  # Generates SQLC vode
	sqlc generate 

.PHONY: test
test:  # Run Unit Tests
	go test -v -cover ./...

.PHONY: start-dev
start-dev:  # Run the Gin Dev Server
	go run main.go