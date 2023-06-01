package types

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
}

func NewUser(email string, passwordHash string) (User, error) {
	user := User{
		Id:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}

	if err := ValidateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func ValidateUser(user User) error {
	errMsg := "Invalid user"

	if user.Id.String() == "" {
		return fmt.Errorf(errMsg + " (empty id)")
	}

	if user.Email == "" {
		return fmt.Errorf(errMsg + " (empty email)")
	}

	if user.PasswordHash == "" {
		return fmt.Errorf(errMsg + " (empty password hash)")
	}

	return nil
}
