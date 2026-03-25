FROM golang:1.26-alpine

ENV ENVIRONMENT=DEV

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go install github.com/cespare/reflex@v0.3.2 \
    && go mod download

ENTRYPOINT ["reflex", "-r", "\\.go$", "-s", "--", "sh", "-c", "go run ./cmd/server"]
