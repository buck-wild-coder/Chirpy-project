package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	cfg := &apiConfig{}
	port := "8080"
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: cfg.routes(),
	}
	log.Printf("Server running on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
