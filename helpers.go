package main

import (
	"encoding/json"
	"log"
	"net/http"
	"unicode/utf8"
)

type Validate struct {
	Body string `json:"body"`
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	response := struct {
		Error string `json:"error"`
	}{
		Error: msg,
	}
	jsonData, _ := json.Marshal(response)
	w.WriteHeader(code)
	w.Write(jsonData)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("can't parse go type to json")
	}
	w.WriteHeader(code)
	w.Write(jsonData)
}

func (v *Validate) check(valid Validate) (interface{}, int) {
	var code int
	var new_data struct {
		Valid bool `json:"valid"`
	}
	if utf8.RuneCountInString(valid.Body) > 140 {
		new_data.Valid = false
		code = 400
	} else {
		new_data.Valid = true
		code = 200
	}
	return new_data, code
}
