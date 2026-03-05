package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/buck-wild-coder/Chirpy-project/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type response struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	params := User{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	hash, err := cfg.db.GetHash(r.Context(), params.Email)
	if err != nil {
		return
	}
	pass, err := auth.CheckPasswordHash(params.Password, hash)
	if !pass {
		respondWithError(w, http.StatusUnauthorized, "401 Unauthorized", fmt.Errorf("401 Unauthorized"))
	}
	respondWithJSON(w, http.StatusOK, User{
		ID:        params.ID,
		CreatedAt: params.CreatedAt,
		UpdatedAt: params.UpdatedAt,
		Email:     params.Email,
		Password:  params.Password,
	})
}
