package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (cfg *apiConfig) routes() *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})

	fs := http.FileServer(http.Dir("."))
	r.PathPrefix("/app").Handler(cfg.middlewareMetricsInc(http.StripPrefix("/app", fs)))

	r.HandleFunc("/api/healthz", healthz).Methods("GET")
	r.HandleFunc("/admin/metrics", cfg.metrics).Methods("GET")
	r.HandleFunc("/admin/reset", cfg.reset).Methods("POST")
	return r
}
