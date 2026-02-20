package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	port := "8080"
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	mux.Handle("/", http.FileServer(http.Dir(".")))
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe error", err)
	}
}
