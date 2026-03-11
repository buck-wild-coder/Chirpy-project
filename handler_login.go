package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/buck-wild-coder/Chirpy-project/internal/auth"
	"github.com/buck-wild-coder/Chirpy-project/internal/database"
	"github.com/google/uuid"
)

type logins struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	Ischirpyred  bool      `json:"is_chirpy_red"`
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

	db, err := cfg.db.GetHash(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "401 Unauthorized", fmt.Errorf("401 Unauthorized"))
		return
	}
	pass, err := auth.CheckPasswordHash(params.Password, db.HashedPassword)
	if !pass {
		respondWithError(w, http.StatusUnauthorized, "401 Unauthorized", fmt.Errorf("401 Unauthorized"))
	}

	tokenString, err := auth.MakeJWT(db.ID, cfg.secret, time.Hour)
	token := auth.MakeRefreshToken()
	arg := database.CreateRefreshTokenParams{
		Token:     token,
		UserID:    db.ID,
		ExpiresAt: time.Now().UTC().AddDate(0, 0, 60),
	}
	returnToken, err := cfg.db.CreateRefreshToken(r.Context(), arg)
	if err != nil {
		return
	}

	login := logins{
		ID:           db.ID,
		Email:        params.Email,
		CreatedAt:    db.CreatedAt,
		UpdatedAt:    db.UpdatedAt,
		Token:        tokenString,
		RefreshToken: returnToken.Token,
		Ischirpyred:  db.IsChirpyRed,
	}

	respondWithJSON(w, http.StatusOK, login)
}
