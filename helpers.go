package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"
)

type Validate struct {
	Body string `json:"body"`
}

func respondWithError(w http.ResponseWriter, status int, msg string) {
	payload := struct {
		Error string `json:"error"`
	}{
		Error: msg,
	}
	data, _ := json.Marshal(payload)
	w.WriteHeader(status)
	w.Write(data)
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("can't parse go type to json")
	}
	w.WriteHeader(status)
	w.Write(data)
}

func (v *Validate) check(input Validate) (Validate, int) {
	if utf8.RuneCountInString(input.Body) > 140 {
		return Validate{}, http.StatusBadRequest
	}
	input.Body = v.filter(input.Body)
	return input, http.StatusOK
}

func (v *Validate) filter(sentence string) string {
	words := strings.Split(sentence, " ")
	var b strings.Builder
	for _, w := range words {
		switch strings.ToLower(w) {
		case "kerfuffle", "sharbert", "fornax":
			w = "****"
		}
		b.WriteString(w + " ")
	}
	return strings.TrimSpace(b.String())
}
