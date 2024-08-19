#!/usr/bin/env bash

fetchENV() {
	cat ./.env | grep $1 | awk -F "=" '{print $2}'
}

PG_USER=$(fetchENV POSTGRES_USER) \
PG_PASS=$(fetchENV POSTGRES_PASSWORD) \
PG_HOST=$(fetchENV POSTGRES_HOST) \
PG_PORT=$(fetchENV POSTGRES_PORT) \
PG_DB=$(fetchENV POSTGRES_DB_NAME) \
	sqlc generate --no-remote


