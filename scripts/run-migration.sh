#!/usr/bin/env sh
set -e

PATH_TO_MIGRATIONS=${1}

>&2 echo "Running migration ..."
migrate -path=${PATH_TO_MIGRATIONS} -database=postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DB?sslmode=disable up
