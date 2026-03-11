package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	value := headers.Get("Authorization")
	if value == "" {
		return "", errors.New("Not Found")
	}

	list := strings.Split(value, " ")
	return list[1], nil
}

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Not Found")
	}
	token := strings.Split(authHeader, " ")

	return token[1], nil
}
