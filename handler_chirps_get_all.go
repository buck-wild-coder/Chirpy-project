package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsGetAll(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("author_id")
	if s != "" {
		userID, err := uuid.Parse(s)
		if err != nil {
			return
		}
		dbChirps, err := cfg.db.GetChirpsAuthor(r.Context(), userID)
		out := make([]Chirp, 0, len(dbChirps))
		for _, c := range dbChirps {
			out = append(out, Chirp{
				ID:        c.ID,
				CreatedAt: c.CreatedAt,
				UpdatedAt: c.UpdatedAt,
				Body:      c.Body,
				UserID:    c.UserID,
			})
		}
		respondWithJSON(w, http.StatusOK, out)
		return
	}
	dbChirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}
	out := make([]Chirp, 0, len(dbChirps))
	for _, c := range dbChirps {
		out = append(out, Chirp{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			Body:      c.Body,
			UserID:    c.UserID,
		})
	}
	respondWithJSON(w, http.StatusOK, out)
}
