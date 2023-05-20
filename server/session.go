package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
)

const SESSION_TOKEN_COOKIE string = "session-token"

func CreateSessionToken() (string, error) {
	bytes := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return "", err
	}

	sessionToken := base64.URLEncoding.EncodeToString(bytes)

	return sessionToken, nil
}

func SetSessionTokenCookie(w http.ResponseWriter, token string) {
	SetCookie(w, SESSION_TOKEN_COOKIE, token)
}

func GetSessionTokenCookie(r *http.Request) (*http.Cookie, error) {
	return GetCookie(r, SESSION_TOKEN_COOKIE)
}
