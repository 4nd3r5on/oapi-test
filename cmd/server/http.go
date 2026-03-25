package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	apiimpl "github.com/4nd3rs0n/oapi-test/internal/api"
	"github.com/4nd3rs0n/oapi-test/internal/app"
	"github.com/4nd3rs0n/oapi-test/internal/app/auth"
	pkgapi "github.com/4nd3rs0n/oapi-test/pkg/api"
	"github.com/4nd3rs0n/oapi-test/pkg/errs"
	"github.com/rs/cors"
)

type httpErrorHandler struct {
	logger *slog.Logger
}

func (h *httpErrorHandler) handle(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		return
	}
	slogArgs := make([]any, 0)

	if cliendData, ok := auth.CtxGetClientData(ctx); ok {
		slogArgs = append(slogArgs, "auth_method", cliendData.Method.String())
	}
	if userID, ok := auth.CtxGetUserID(ctx); ok {
		slogArgs = append(slogArgs, "user_id", userID.String())
	}

	// TODO: Maybe Add some headers to logs?
	errs.HandleHTTP(ctx, h.logger, w, err, slogArgs...)
}

func httpAPI(appLayer *app.App, authOpts Auth, logger *slog.Logger, addr string) {
	mux := http.NewServeMux()

	apiHandler, err := apiimpl.NewAPIHandler(appLayer, logger)
	if err != nil {
		log.Fatalf("failed to create new API handler: %v", err)
	}
	securityHandler := &apiimpl.SecurityHandler{
		TMB: authOpts.TMB,
	}

	errHandler := &httpErrorHandler{logger: logger}
	handler, err := pkgapi.NewServer(
		apiHandler,
		securityHandler,
		pkgapi.WithErrorHandler(errHandler.handle),
	)
	if err != nil {
		log.Fatalf("failed to initlize the server: %v", err)
	}

	apiPrefix := "/api/v1"
	mux.Handle(apiPrefix, http.StripPrefix(apiPrefix, handler))
	httpHandler := cors.AllowAll().Handler(mux)

	logger.Info("running server", "addr", addr)
	if err := http.ListenAndServe(addr, httpHandler); err != nil {
		log.Fatalf("error running the server: %v", err)
	}
}
