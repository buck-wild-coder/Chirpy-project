package main

import (
	"errors"
	"net/http"

	"github.com/buck-wild-coder/Chirpy-project/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handerChirpDelete(w http.ResponseWriter, r *http.Request) {
	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, tokenFailure, errors.New(tokenFailure))
		return
	}
	UserID, err := auth.ValidateJWT(tokenString, cfg.secret)
	if err != nil {
		respondWithError(w, 401, unauthorizedAccess, errors.New(unauthorizedAccess))
		return
	}
	chirp := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirp)
	if err != nil {
		return
	}
	db, err := cfg.db.GetAChirp(r.Context(), chirpID)
	if err != nil {
		return
	}

	if db.UserID != UserID {
		respondWithError(w, 403, "Unauthorized chirp", errors.New("unsuthorized"))
		return
	}
	err = cfg.db.DeleteChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, 403, "Error deleting chirp", err)
	}
	respondWithJSON(w, 204, "deleted Successfuly")
}
