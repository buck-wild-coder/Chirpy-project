package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (cfg *apiConfig) metrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	hitCount := cfg.fileserverHits.Load()
	template := `<html>
			<body>
				<h1>Welcome, Chirpy Admin</h1>
				<p>Chirpy has been visited %d times!</p>
			</body>
			</html>`
	html := fmt.Sprintf(template, hitCount)
	fmt.Fprint(w, html)
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

func validate_chirp(w http.ResponseWriter, r *http.Request) {
	var payload Validate
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, 400, "Bad request, try again.")
	}
	json.Unmarshal(bodyBytes, &payload)
	cleaned, statusCode := payload.check(payload)
	response := struct {
		CleanedBody string `json:"cleaned_body"`
	}{
		CleanedBody: cleaned.Body,
	}
	respondWithJSON(w, statusCode, response)
}
