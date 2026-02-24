package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/buck-wild-coder/Chirpy-project/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	dbQueries      *database.Queries
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Println(err)
	}
	dbQueries := database.New(db)

	cfg := &apiConfig{
		dbQueries: dbQueries,
	}

	port := "8080"
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: cfg.routes(),
	}
	log.Printf("Server running on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
