package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) metrics(w http.ResponseWriter, r *http.Request) {
	value := cfg.fileserverHits.Load()
	fmt.Fprintf(w, "Hits: %d", value)
}

func (cfg *apiConfig) reset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.Write([]byte("Reset"))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
