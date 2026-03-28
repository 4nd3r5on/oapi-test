package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/4nd3rs0n/oapi-test/internal/app"
	"github.com/4nd3rs0n/oapi-test/internal/app/auth"
	"github.com/4nd3rs0n/oapi-test/internal/config"
	"github.com/4nd3rs0n/oapi-test/internal/infra/postgres"
	"github.com/4nd3rs0n/oapi-test/internal/infra/redis"
)

type Auth struct {
	TMA, TMB, Session auth.VerifierFunc
}

func main() {
	// Config
	ctx := context.Background()
	env := config.GetEnvironment()
	if env == config.EnvironmentUnknown {
		log.Fatalf(
			"missconfigured or missing required enviromnment variable %s",
			config.EnvEnvironment,
		)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: env.LogLevel(),
	}))
	slog.SetDefault(logger)

	tgBotToken := os.Getenv(config.EnvTGBotToken)
	dbURL := getRequiredEnv(config.EnvDBURL)
	cacheURL := getRequiredEnv(config.EnvCacheURL)
	apiAddr := getEnv(config.EnvAPIAddr, ":9090")

	// Connecting to services

	slog.Debug("connecting DB", "url", dbURL)
	db, err := postgres.NewDB(ctx, dbURL)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer db.Close()

	// TODO: Add S3 for profile pictures and attachments
	// slog.Debug("connecting S3 storage", "url", s3URL)
	// storage, err := minio.NewStorage(ctx, s3URL)
	// if err != nil {
	// 	log.Fatalf("storage connection failed: %v", err)
	// }

	slog.Debug("connecting cache", "url", cacheURL)
	cache, err := redis.NewCache(ctx, cacheURL)
	if err != nil {
		log.Fatalf("cache connection failed: %v", err)
	}
	defer cache.Close()

	// Initializing the app

	appLayer, err := app.New(db, cache, nil)
	if err != nil {
		log.Fatalf("error initializing application layer: %v", err)
	}

	var authOpts Auth
	if env == config.EnvironmentDev || env == config.EnvironmentTest {
		authOpts.TMB = (&auth.TMB{}).Verify
	}
	if tgBotToken != "" {
		authOpts.TMA = (&auth.TMA{BotToken: tgBotToken}).Verify
	}
	// TODO: Sessions auth

	// Running the app

	var wg sync.WaitGroup
	wg.Go(func() {
		httpAPI(appLayer, authOpts, logger, apiAddr)
	})
	wg.Wait()
}
