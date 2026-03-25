package main

import (
	"log/slog"
	"os"
)

func getRequiredEnv(name string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		slog.Error("Required environment variable wasn't set", "name", name)
		os.Exit(1)
	}
	return val
}

func getEnv(name, fallback string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		return fallback
	}
	return val
}

