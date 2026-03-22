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
	mux := http.NewServeMux()

	api := &apiimpl.API{
		Ctx:    ctx,
		App:    nil,
		Logger: logger,
	}
	handler := pkgapi.HandlerFromMux(api, mux)

	logger.Info("running server", "url", url)
	if err := http.ListenAndServe(url, handler); err != nil {
		log.Fatalf("error running the server: %v", err)
	}
}

func main() {
	ctx := context.Background()
	httpAPI(ctx, slog.Default())
}
