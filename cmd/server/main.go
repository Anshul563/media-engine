package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/Anshul563/media-engine/internal/config"
	"github.com/Anshul563/media-engine/internal/logger"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.App.Env)

	slog.SetDefault(log)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		slog.Info(
			"health check",
			"method", r.Method,
			"path", r.URL.Path,
		)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte(`{"status":"ok"}`))
		if err != nil {
			slog.Error(
				"failed to write health response",
				"error", err,
			)
		}
	})

	addr := fmt.Sprintf(":%d", cfg.App.Port)

	slog.Info(
		"media engine starting",
		"name", cfg.App.Name,
		"environment", cfg.App.Env,
		"address", addr,
	)

	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error(
			"server stopped",
			"error", err,
		)

		os.Exit(1)
	}
}
