# OpenAPI playground

My learning repo to learn about openapi, how to work with it, hot to integrate it into my flow.

## Configure

Create `.env` file and fill it out with the content
```sh
cp example.env .env
# generate passwords
pw=$(openssl rand -hex 24) && sed -i "s/^APP_PG_PASS=.*/APP_PG_PASS=$pw/" .env
pw=$(openssl rand -hex 24) && sed -i "s/^PG_PASS=.*/PG_PASS=$pw/" .env
pw=$(openssl rand -hex 24) && sed -i "s/^REDIS_PASSWORD=.*/REDIS_PASSWORD=$pw/" .env

# Set S3_URL
nvim .env
```

## Dependencies

- Go
- Docker & Docker compose
- S3 (for Local Development seaweedfs)

### Development

- [SQLc](https://sqlc.dev)
- [Golang Migrate](https://github.com/golang-migrate/migrate)
- [GolangCI Lint](https://golangci-lint.run/)

## Development

```sh
# Generate code from openapi docs
# 1. Bundle
npx @redocly/cli bundle ./pkg/api/openapi/api.yml --ext yml -o ./pkg/api/api.yml
# 2. Generate
go generate ./...

# Generate db package
sqlc generate

# Apply migrations (only dev container)
migrate -path ./migrations -database $DB_URL up

# Test
go test -v ./...

# Lint
golangci-lint run --fix
```

## Run dev

Running code with hot reload

```sh
# Up the DB
docker compose -f ./compose.dev.yml up db -d
# Apply migrations
set -a && source .env && set +a
export DB_URL=postgres://$APP_PG_USER:$APP_PG_PASS@127.0.0.1:5432/$APP_PG_DB?sslmode=disable
migrate -path pkg/db/migrations -database $DB_URL up
# Run everything
docker compose -f ./compose.dev.yml up --build
```
