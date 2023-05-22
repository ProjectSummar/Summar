package postgres

import (
	"database/sql"
	"fmt"
	"summar/server/types"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	Db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=123 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		Db: db,
	}, nil
}

// Operations

func (s *PostgresStore) CreateSession(session *types.Session) error {
	query := `insert into sessions
	(token, user_id, expires_at)
	values ($1, $2, $3)
	`

	res, err := s.Db.Query(
		query,
		session.Token,
		session.UserId,
		session.ExpiresAt,
	)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)
	return nil
}

func (s *PostgresStore) GetSession(token string) (*types.Session, error) {
	query := "SELECT * FROM sessions WHERE token = $1"

	rows, err := s.Db.Query(query, token)
	if err != nil {
		return nil, err
	}

	var session types.Session
	for rows.Next() {
		if err = rows.Scan(
			&session.Token,
			&session.UserId,
			&session.ExpiresAt,
		); err != nil {
			return nil, err
		}
	}

	fmt.Printf("%+v\n", session)
	return &session, nil
}

func (s *PostgresStore) CreateUser(user *types.User) error {
	query := `insert into users
	(id, email, password_hash, created_at)
	values ($1, $2, $3, $4)
	`

	res, err := s.Db.Query(
		query,
		user.Id,
		user.Email,
		user.PasswordHash,
		user.CreatedAt,
	)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)
	return nil
}

func (s *PostgresStore) GetUser(id uuid.UUID) (*types.User, error) {
	query := "SELECT * FROM users WHERE id = $1"

	rows, err := s.Db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var user types.User
	for rows.Next() {
		if err = rows.Scan(
			&user.Id,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
	}

	fmt.Printf("%+v\n", user)
	return &user, nil
}

func (s *PostgresStore) GetUserByEmail(email string) (*types.User, error) {
	query := "SELECT * FROM users WHERE email = $1"

	rows, err := s.Db.Query(query, email)
	if err != nil {
		return nil, err
	}

	var user types.User
	for rows.Next() {
		if err = rows.Scan(
			&user.Id,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
	}

	fmt.Printf("%+v\n", user)
	return &user, nil
}

func (s *PostgresStore) UpdateUser(user *types.User) error {
	return nil
}

func (s *PostgresStore) DeleteUser(userId uuid.UUID) error {
	return nil
}

// Initialisation

func (s *PostgresStore) Init() error {
	if err := s.CreateUsersTable(); err != nil {
		return err
	}

	if err := s.CreateSessionsTable(); err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) CreateUsersTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(36),
		email VARCHAR(50),
		password_hash VARCHAR(60),
		created_at TIMESTAMP,
		PRIMARY KEY(id),
		UNIQUE(email)
	)`

	_, err := s.Db.Exec(query)
	return err
}

func (s *PostgresStore) CreateSessionsTable() error {
	query := `CREATE TABLE IF NOT EXISTS sessions (
	token VARCHAR(44),
	user_id VARCHAR(36),
	expires_at TIMESTAMP,
	PRIMARY KEY(token),
	CONSTRAINT fk_user
		FOREIGN KEY(user_id)
			REFERENCES users(id)
			ON DELETE CASCADE
	)`

	_, err := s.Db.Exec(query)
	return err
}

func (s *PostgresStore) Clear() error {
	s.Db.Exec("DROP TABLE IF EXISTS sessions")
	s.Db.Exec("DROP TABLE IF EXISTS users")

	return nil
}
