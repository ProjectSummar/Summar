package stores

import (
	"fmt"
	"summar/server/types"

	"github.com/google/uuid"
)

type MockStore struct {
	Users    []*types.User
	Sessions []*types.Session
}

func NewMockStore() *MockStore {
	return &MockStore{
		Users:    []*types.User{},
		Sessions: []*types.Session{},
	}
}

func (s *MockStore) CreateSession(session *types.Session) error {
	s.Sessions = append(s.Sessions, session)
	return nil
}

func (s *MockStore) GetSession(token string) (*types.Session, error) {
	for _, session := range s.Sessions {
		if session.Token == token {
			return session, nil
		}
	}

	return nil, fmt.Errorf("Session not found")
}

func (s *MockStore) CreateUser(user *types.User) error {
	s.Users = append(s.Users, user)
	return nil
}

func (s *MockStore) GetUser(id uuid.UUID) (*types.User, error) {
	for _, user := range s.Users {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, fmt.Errorf("User not found")
}

func (s *MockStore) GetUserByEmail(email string) (*types.User, error) {
	for _, user := range s.Users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, fmt.Errorf("User not found")
}
