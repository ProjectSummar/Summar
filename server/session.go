package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"time"
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

type InvalidSessionError struct{}

func (e *InvalidSessionError) Error() string {
	return "Invalid session token found (expired)"
}

func ValidateSessionExpiry(session *Session) error {
	invalid := session.ExpiresAt.Compare(time.Now()) < 0
	if invalid {
		return &InvalidSessionError{}
	} else {
		return nil
	}
}
