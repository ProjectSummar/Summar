package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type Bookmark struct {
	ID      string    `json:"id"`
	UserID  uuid.UUID `json:"userId"`
	Url     string    `json:"url"`
	Summary string    `json:"summary"`
}
