package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/buck-wild-coder/Chirpy-project/internal/auth"
	"github.com/google/uuid"
)

type logins struct {
	ID                 uuid.UUID `json:"id"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Expires_in_seconds int       `json:"expires_in_seconds"`
	Token              string    `json:"token"`
}

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type response struct {
		logins
	}

	decoder := json.NewDecoder(r.Body)
	params := logins{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Expires_in_seconds == 0 {
		params.Expires_in_seconds = 1
	}
	db, err := cfg.db.GetHash(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "401 Unauthorized", fmt.Errorf("401 Unauthorized"))
		return
	}
	pass, err := auth.CheckPasswordHash(params.Password, db.HashedPassword)
	if !pass {
		respondWithError(w, http.StatusUnauthorized, "401 Unauthorized", fmt.Errorf("401 Unauthorized"))
	}

	login := logins{
		ID:                 db.ID,
		Email:              params.Email,
		CreatedAt:          db.CreatedAt,
		UpdatedAt:          db.UpdatedAt,
		Expires_in_seconds: params.Expires_in_seconds,
	}
	duration := time.Duration(login.Expires_in_seconds) * time.Second
	tokenString, err := auth.MakeJWT(login.ID, cfg.secret, duration)
	login.Token = tokenString

	respondWithJSON(w, http.StatusOK, login)
}
