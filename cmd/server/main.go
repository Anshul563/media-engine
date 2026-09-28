package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Anshul563/media-engine/internal/config"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	addr := fmt.Sprintf(":%d", cfg.App.Port)

	log.Printf(
		"%s running in %s mode on http://localhost%s",
		cfg.App.Name,
		cfg.App.Env,
		addr,
	)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}