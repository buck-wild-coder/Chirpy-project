package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/buck-wild-coder/Chirpy-project/internal/auth"
	"github.com/buck-wild-coder/Chirpy-project/internal/database"
)

func (cfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
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
	var Input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&Input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "CAn't decode", errors.New("can't decode"))
		return
	}
	pass, err := auth.HashPassword(Input.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Can't hash the password", err)
		return
	}
	user, err := cfg.db.UpdateUser(r.Context(), database.UpdateUserParams{
		Email:          Input.Email,
		HashedPassword: pass,
		ID:             UserID,
	})

	respondWithJSON(w, http.StatusOK, User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Email:       user.Email,
		Password:    user.HashedPassword,
		Ischirpyred: user.IsChirpyRed,
	})
}
