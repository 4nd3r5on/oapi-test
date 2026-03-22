package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	apiimpl "github.com/4nd3rs0n/oapi-test/internal/api"
	pkgapi "github.com/4nd3rs0n/oapi-test/pkg/api"
)

func httpAPI(ctx context.Context, logger *slog.Logger) {
	url := ":9090"
	apiHandler, err := apiimpl.NewAPIHandler(
		ctx, nil, logger,
	)
	if err != nil {
		log.Fatalf("failed to create new API handler: %w", err)
	}
	handler, err := pkgapi.NewServer(apiHandler, nil)
	if err != nil {
		log.Fatalf("failed to initlize the server: %w", err)
	}

	logger.Info("running server", "url", url)
	if err := http.ListenAndServe(url, handler); err != nil {
		log.Fatalf("error running the server: %v", err)
	}
}

func main() {
	ctx := context.Background()
	httpAPI(ctx, slog.Default())
}
