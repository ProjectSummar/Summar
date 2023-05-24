package stores

import (
	"database/sql"
	"fmt"
	"summar/server/types"
	"summar/server/utils"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	Db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=123 sslmode=disable"

	var db *sql.DB
	var retryErr error

	for retries := 10; retries > 0; retries-- {
		db, retryErr = sql.Open("postgres", connStr)
		if retryErr != nil {
			fmt.Println(retryErr)
			fmt.Println("Retries left:", retries)
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	if retryErr != nil {
		return nil, retryErr
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

	_, err := s.Db.Query(
		query,
		session.Token,
		session.UserId,
		session.ExpiresAt,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Session created\n%+v\n", utils.JSONMarshal(session))
	return nil
}

func (s *PostgresStore) GetSession(token string) (*types.Session, error) {
	rows, err := s.Db.Query("SELECT * FROM sessions WHERE token = $1", token)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return ScanSessionRow(rows)
	}

	return nil, fmt.Errorf("Session not found")
}

func (s *PostgresStore) CreateUser(user *types.User) error {
	query := `insert into users
	(id, email, password_hash, created_at)
	values ($1, $2, $3, $4)
	`

	_, err := s.Db.Query(
		query,
		user.Id,
		user.Email,
		user.PasswordHash,
		user.CreatedAt,
	)
	if err != nil {
		return err
	}

	fmt.Printf("User created\n%+v\n", utils.JSONMarshal(user))
	return nil
}

func (s *PostgresStore) GetUser(id uuid.UUID) (*types.User, error) {
	rows, err := s.Db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return ScanUserRow(rows)
	}

	return nil, fmt.Errorf("User not found")
}

func (s *PostgresStore) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.Db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return ScanUserRow(rows)
	}

	return nil, fmt.Errorf("User not found")
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

func (s *PostgresStore) Clear() {
	s.Db.Exec("DROP TABLE IF EXISTS sessions")
	s.Db.Exec("DROP TABLE IF EXISTS users")
}

// Helpers

func ScanSessionRow(rows *sql.Rows) (*types.Session, error) {
	var session types.Session
	err := rows.Scan(
		&session.Token,
		&session.UserId,
		&session.ExpiresAt,
	)

	return &session, err
}

func ScanUserRow(rows *sql.Rows) (*types.User, error) {
	var user types.User
	err := rows.Scan(
		&user.Id,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	return &user, err
}