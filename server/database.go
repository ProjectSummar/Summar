package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type DB interface {
	Clear() error

	// session
	CreateSession(*Session) error
	GetSessionByToken(string) (*Session, error)

	// user
	CreateUser(*User) error
	// DeleteUser(string) error
	// UpdateUser(*User) error
	GetUserByEmail(string) (*User, error)
	GetUserById(uuid.UUID) (*User, error)

	// bookmark
	GetBookmarksByUserId(uuid.UUID) ([]Bookmark, error)
}

type PostgresDB struct {
	Db *sql.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	connStr := "user=postgres dbname=postgres password=123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{
		Db: db,
	}, nil
}

func (db *PostgresDB) CreateSession(session *Session) error {
	query := `insert into sessions
	(token, user_id, expires_at)
	values ($1, $2, $3)
	`

	res, err := db.Db.Query(
		query,
		session.Token,
		session.UserID,
		session.ExpiresAt,
	)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)
	return nil
}

func (db *PostgresDB) GetSessionByToken(token string) (*Session, error) {
	query := "SELECT * FROM sessions WHERE token = $1"

	rows, err := db.Db.Query(query, token)
	if err != nil {
		return nil, err
	}

	var session Session
	for rows.Next() {
		if err = rows.Scan(
			&session.Token,
			&session.UserID,
			&session.ExpiresAt,
		); err != nil {
			return nil, err
		}
	}

	fmt.Printf("%+v\n", session)
	return &session, nil
}

func (db *PostgresDB) CreateUser(user *User) error {
	query := `insert into users
	(id, email, password_hash, created_at)
	values ($1, $2, $3, $4)
	`

	res, err := db.Db.Query(
		query,
		user.ID,
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

func (db *PostgresDB) GetUserByEmail(email string) (*User, error) {
	query := "SELECT * FROM users WHERE email = $1"

	rows, err := db.Db.Query(query, email)
	if err != nil {
		return nil, err
	}

	var user User
	for rows.Next() {
		if err = rows.Scan(
			&user.ID,
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

func (db *PostgresDB) GetUserById(id uuid.UUID) (*User, error) {
	query := "SELECT * FROM users WHERE id = $1"

	rows, err := db.Db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var user User
	for rows.Next() {
		if err = rows.Scan(
			&user.ID,
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

func (db *PostgresDB) GetBookmarksByUserId(userId uuid.UUID) ([]Bookmark, error) {
	// TODO
	return nil, nil
}

// Initialisation

func (db *PostgresDB) Init() error {
	if err := db.CreateUsersTable(); err != nil {
		return err
	}

	if err := db.CreateSessionsTable(); err != nil {
		return err
	}

	return nil
}

func (db *PostgresDB) CreateUsersTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(36),
		email VARCHAR(50),
		password_hash VARCHAR(60),
		created_at TIMESTAMP,
		PRIMARY KEY(id),
		UNIQUE(email)
	)`

	_, err := db.Db.Exec(query)
	return err
}

func (db *PostgresDB) CreateSessionsTable() error {
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

	_, err := db.Db.Exec(query)
	return err
}

func (db *PostgresDB) Clear() error {
	db.Db.Exec("DROP TABLE IF EXISTS sessions")
	db.Db.Exec("DROP TABLE IF EXISTS users")

	return nil
}
