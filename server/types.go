package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
}

func NewUser(email string, passwordHash string) User {
	return User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}
}

type Session struct {
	Token     string    `json:"token"`
	UserID    uuid.UUID `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func NewSession(token string, userId uuid.UUID) Session {
	return Session{
		Token:     token,
		UserID:    userId,
		ExpiresAt: time.Now().Add(time.Second * time.Duration(SESSION_EXPIRATION_SECONDS)),
	}
}

type Bookmark struct {
	ID      string    `json:"id"`
	UserID  uuid.UUID `json:"userId"`
	Url     string    `json:"url"`
	Summary string    `json:"summary"`
}
