#!/usr/bin/env bash

export PG_URL="postgres://user:myAwEsOm3pa55@w0rd@localhost:5432/db"
export PG_HOST="localhost"
export PG_USERNAME="postgres"
export PG_PASSWORD="postgres"
export PG_DBNAME="postgres"
export PG_PORT="5432"
export PG_SSLMODE="disable"

echo $PG_URL
air -c air.toml
