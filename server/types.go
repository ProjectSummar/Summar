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

type Session struct {
	Token     string    `json:"token"`
	UserID    uuid.UUID `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type Bookmark struct {
	ID      string    `json:"id"`
	UserID  uuid.UUID `json:"userId"`
	Url     string    `json:"url"`
	Summary string    `json:"summary"`
}
