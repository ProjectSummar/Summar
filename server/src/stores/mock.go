package stores

import (
	"fmt"
	"summar/server/types"

	"github.com/google/uuid"
)

type MockStore struct {
	Users     []types.User
	Sessions  []types.Session
	Bookmarks []types.Bookmark
}

func NewMockStore() *MockStore {
	return &MockStore{
		Users:    []types.User{},
		Sessions: []types.Session{},
	}
}

func (s *MockStore) CreateSession(session types.Session) error {
	s.Sessions = append(s.Sessions, session)

	return nil
}

func (s *MockStore) GetSession(token string) (types.Session, error) {
	for _, session := range s.Sessions {
		if session.Token == token {
			return session, nil
		}
	}

	return types.Session{}, fmt.Errorf("Session not found")
}

func (s *MockStore) CreateUser(user types.User) error {
	s.Users = append(s.Users, user)

	return nil
}

func (s *MockStore) GetUser(id uuid.UUID) (types.User, error) {
	for _, user := range s.Users {
		if user.Id == id {
			return user, nil
		}
	}

	return types.User{}, fmt.Errorf("User not found")
}

func (s *MockStore) GetUserByEmail(email string) (types.User, error) {
	for _, user := range s.Users {
		if user.Email == email {
			return user, nil
		}
	}

	return types.User{}, fmt.Errorf("User not found")
}

func (s *MockStore) CreateBookmark(bookmark types.Bookmark) error {
	s.Bookmarks = append(s.Bookmarks, bookmark)
	return nil
}

func (s *MockStore) GetBookmark(id uuid.UUID) (types.Bookmark, error) {
	for _, bookmark := range s.Bookmarks {
		if bookmark.Id == id {
			return bookmark, nil
		}
	}

	return types.Bookmark{}, fmt.Errorf("Bookmark not found")
}

func (s *MockStore) GetBookmarksByUserId(userId uuid.UUID) ([]types.Bookmark, error) {
	bookmarks := []types.Bookmark{}

	for _, bookmark := range s.Bookmarks {
		if bookmark.UserId == userId {
			bookmarks = append(bookmarks, bookmark)
		}
	}

	return bookmarks, nil
}

func (s *MockStore) UpdateBookmark(b types.Bookmark) error {
	bookmark, err := s.GetBookmark(b.Id)
	if err != nil {
		return err
	}

	if b.Url != "" {
		bookmark.Url = b.Url
	}

	if b.Summary != "" {
		bookmark.Summary = b.Summary
	}

	return nil
}

func (s *MockStore) DeleteBookmark(id uuid.UUID) error {
	for idx, bookmark := range s.Bookmarks {
		if bookmark.Id == id {
			s.Bookmarks[idx] = s.Bookmarks[len(s.Bookmarks)-1]
			s.Bookmarks = s.Bookmarks[:len(s.Bookmarks)-1]
		}
	}

	return nil
}
