#!/bin/sh

# Will immediately terminate script if got non 0 (error)
set -e

# Reading env file for environment variables
# source ./app.env

echo "Running Migrations"
/app/migrate -path /app/db/migrations -database "$DB_SOURCE" -verbose up
# make migration-up

echo "Starting the app"
exec "$@"