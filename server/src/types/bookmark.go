package types

import (
	"fmt"

	"github.com/google/uuid"
)

type Bookmark struct {
	Id      uuid.UUID `json:"id"`
	UserId  uuid.UUID `json:"userId"`
	Url     string    `json:"url"`
	Title   string    `json:"title"`
	Summary string    `json:"summary"`
}

func NewBookmark(userId uuid.UUID, url string, title string) (Bookmark, error) {
	bookmark := Bookmark{
		Id:     uuid.New(),
		UserId: userId,
		Url:    url,
		Title:  title,
	}

	if err := ValidateBookmark(bookmark); err != nil {
		return Bookmark{}, err
	}

	return bookmark, nil
}

func ValidateBookmark(bookmark Bookmark) error {
	err := "Invalid bookmark"

	if bookmark.Id.String() == "" {
		return fmt.Errorf(err + " (empty id)")
	}

	if bookmark.UserId.String() == "" {
		return fmt.Errorf(err + " (empty user id)")
	}

	if bookmark.Url == "" {
		return fmt.Errorf(err + " (empty url)")
	}

	return nil
}
