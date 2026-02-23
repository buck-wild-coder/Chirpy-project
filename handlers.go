package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (cfg *apiConfig) metrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	value := cfg.fileserverHits.Load()
	body := `<html>
			<body>
				<h1>Welcome, Chirpy Admin</h1>
				<p>Chirpy has been visited %d times!</p>
			</body>
			</html>`
	html := fmt.Sprintf(body, value)
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
	var valid Validate
	data, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, 400, "Bad request, try again.")
	}
	json.Unmarshal(data, &valid)
	new_data, code := valid.check(valid)
	respondWithJSON(w, code, new_data)
}
