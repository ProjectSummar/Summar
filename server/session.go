package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func CreateSessionToken() (string, error) {
	bytes := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return "", err
	}

	sessionToken := base64.URLEncoding.EncodeToString(bytes)

	return sessionToken, nil
}
