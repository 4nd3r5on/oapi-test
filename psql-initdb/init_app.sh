#!/usr/bin/env bash
set -euo pipefail

# Standard Arch-style: Explicit env checks
: "${APP_PG_USER:=app}"
: "${APP_PG_DB:=app}"
: "${APP_PG_PASS:?Error: APP_PG_PASS is not set}"

echo "Running idempotent init for $APP_PG_DB..."

psql -v ON_ERROR_STOP=1 \
  --username "$POSTGRES_USER" \
  -v db_name="$APP_PG_DB" \
  -v db_user="$APP_PG_USER" \
  -v db_pass="$APP_PG_PASS" \
  -f /docker-entrypoint-initdb.d/manual/create-db-and-role.sql
