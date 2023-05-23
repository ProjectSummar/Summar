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

type User struct {
	Id           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
}

func NewUser(email string, passwordHash string) User {
	return User{
		Id:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}
}

type Session struct {
	Token     string    `json:"token"`
	UserId    uuid.UUID `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func NewSession(userId uuid.UUID) Session {
	bytes := make([]byte, 32)
	io.ReadFull(rand.Reader, bytes)
	token := base64.URLEncoding.EncodeToString(bytes)

	return Session{
		Token:     token,
		UserId:    userId,
		ExpiresAt: time.Now().Add(time.Second * time.Duration(constants.SESSION_EXPIRATION_SECONDS)),
	}
}

func VerifySessionExpiry(session *Session) error {
	invalid := session.ExpiresAt.Compare(time.Now()) < 0

	if invalid {
		return fmt.Errorf("Session has expired")
	} else {
		return nil
	}
}

type Bookmark struct {
	Id      uuid.UUID `json:"id"`
	UserId  uuid.UUID `json:"userId"`
	Url     string    `json:"url"`
	Summary string    `json:"summary"`
}
