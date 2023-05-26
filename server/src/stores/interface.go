package stores

import (
	"summar/server/types"

	"github.com/google/uuid"
)

type Store interface {
	// session
	CreateSession(*types.Session) error
	GetSession(string) (*types.Session, error)
	// DeleteSession(string) error

	// user
	CreateUser(*types.User) error
	GetUser(uuid.UUID) (*types.User, error)
	GetUserByEmail(string) (*types.User, error)
	// UpdateUser(*types.User) error
	// DeleteUser(uuid.UUID) error

	// bookmark
	CreateBookmark(*types.Bookmark) error
	GetBookmark(uuid.UUID) (*types.Bookmark, error)
	GetBookmarksByUserId(uuid.UUID) ([]*types.Bookmark, error)
	// UpdateBookmark
	// DeleteBookmark
}
