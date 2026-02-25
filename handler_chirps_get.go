package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("chirpID")
	uid, err := uuid.Parse(id)
	if err != nil {
		respondWithError(w, 404, "Bad Request", err)
		return
	}
	chirp, err := cfg.db.GetAChirp(r.Context(), uid)
	if err != nil {
		respondWithError(w, 404, "Could not find it", err)
		return
	}
	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
