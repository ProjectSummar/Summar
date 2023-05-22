package main

import (
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

func NewSession(token string, userId uuid.UUID) Session {
	return Session{
		Token:     token,
		UserId:    userId,
		ExpiresAt: time.Now().Add(time.Second * time.Duration(SESSION_EXPIRATION_SECONDS)),
	}
}

type Bookmark struct {
	Id      string    `json:"id"`
	UserId  uuid.UUID `json:"userId"`
	Url     string    `json:"url"`
	Summary string    `json:"summary"`
}
