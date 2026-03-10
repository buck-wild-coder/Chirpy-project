package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/buck-wild-coder/Chirpy-project/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		return
	}
	db, err := cfg.db.GetToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Cannot retrieve 401 Unauthorized", errors.New("401 Unauthorized"))
		return
	}
	if db.ExpiresAt.Before(time.Now().UTC()) {
		respondWithError(w, http.StatusUnauthorized, "expired 401 Unauthorized", errors.New("401 Unauthorized"))
		return
	}
	if db.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "revoked 401 Unauthorized", errors.New("401 Unauthorized"))
		return
	}

	str, err := auth.MakeJWT(db.UserID, cfg.secret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "StatusInternalServerError", errors.New("StatusInternalServerError"))
		return
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: str,
	}

	respondWithJSON(w, 200, response)
}
