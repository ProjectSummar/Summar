package types

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"summar/server/constants"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Token     string    `json:"token"`
	UserId    uuid.UUID `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func NewSession(userId uuid.UUID) *Session {
	bytes := make([]byte, 32)
	io.ReadFull(rand.Reader, bytes)
	token := base64.URLEncoding.EncodeToString(bytes)

	return &Session{
		Token:     token,
		UserId:    userId,
		ExpiresAt: time.Now().Add(time.Second * time.Duration(constants.SESSION_EXPIRATION_SECONDS)),
	}
}

func ValidateSession(session *Session) error {
	invalid := session.ExpiresAt.Compare(time.Now()) < 0

	if invalid {
		return fmt.Errorf("Session has expired")
	} else {
		return nil
	}
}
