package main

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/buck-wild-coder/Chirpy-project/internal/auth"
	"github.com/buck-wild-coder/Chirpy-project/internal/database"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "UN", errors.New("un"))
		return
	}
	// refreshToken is the string from the header
	err = cfg.db.RevokeRefreshToken(r.Context(), database.RevokeRefreshTokenParams{
		Token:     refreshToken,
		UpdatedAt: time.Now().UTC(),
		RevokedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't revoke token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
