package main

import (
	"context"
	"log/slog"
	"net/http"

	apiimpl "github.com/4nd3rs0n/oapi-test/internal/api"
	pkgapi "github.com/4nd3rs0n/oapi-test/pkg/api"
)

func httpAPI(ctx context.Context) {
	url := ":9090"
	mux := http.NewServeMux()

	api := apiimpl.NewAPI(ctx)
	handler := pkgapi.HandlerFromMux(api, mux)

	slog.Info("Running server", "url", url)
	http.ListenAndServe(url, handler)
}

func main() {
	ctx := context.Background()
	httpAPI(ctx)
}
