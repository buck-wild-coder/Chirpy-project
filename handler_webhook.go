package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/buck-wild-coder/Chirpy-project/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerWebhook(w http.ResponseWriter, r *http.Request) {
	api, err := auth.GetAPIKey(r.Header)
	if api != cfg.PolkaKey {
		respondWithError(w, 401, "Unauthenticated api", errors.New("Unauthenticated api"))
		return
	}
	type user_id struct {
		UserID string `json:"user_id"`
	}
	var Input struct {
		Event string  `json:"event"`
		Data  user_id `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&Input)
	if err != nil {
		return
	}

	userid, err := uuid.Parse(Input.Data.UserID)
	if err != nil {
		return
	}

	if Input.Event == "user.upgraded" {
		err = cfg.db.UpgradeToRed(r.Context(), userid)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Server error", errors.New("Server error"))
			return
		}
	}
	respondWithJSON(w, 204, "")
}
