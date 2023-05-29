package stores

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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
	sleep := time.Second

	for retries := 10; retries > 0; retries-- {
		connStr := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		)

		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Println(err)
			log.Println("Retries left:", retries)
			sleep *= 2
			time.Sleep(sleep)
			continue
		}

		if err := db.Ping(); err != nil {
			log.Println(err)
			log.Println("Retries left:", retries)
			sleep *= 2
			time.Sleep(sleep)
			continue
		}

		return &PostgresStore{
			Db: db,
		}, nil
	}

	return nil, fmt.Errorf("Cannot connect to Postgres")
}

// Session operations

func (s *PostgresStore) CreateSession(session types.Session) error {
	query := `
	insert into sessions
	(token, user_id, expires_at)
	values ($1, $2, $3)
	`

	_, err := s.Db.Exec(
		query,
		session.Token,
		session.UserId,
		session.ExpiresAt,
	)
	if err != nil {
		return err
	}

	log.Printf("Session created\n%+v\n", utils.JSONMarshalIndent(session))
	return nil
}

func (s *PostgresStore) GetSession(token string) (types.Session, error) {
	rows, err := s.Db.Query("SELECT * FROM sessions WHERE token = $1", token)
	if err != nil {
		return types.Session{}, err
	}

	for rows.Next() {
		return ScanSessionRow(rows)
	}

	return types.Session{}, fmt.Errorf("Session not found")
}

// User operations

func (s *PostgresStore) CreateUser(user types.User) error {
	query := `
	insert into users
	(id, email, password_hash, created_at)
	values ($1, $2, $3, $4)
	`

	_, err := s.Db.Exec(
		query,
		user.Id,
		user.Email,
		user.PasswordHash,
		user.CreatedAt,
	)
	if err != nil {
		return err
	}

	log.Printf("User created\n%+v\n", utils.JSONMarshalIndent(user))
	return nil
}

func (s *PostgresStore) GetUser(id uuid.UUID) (types.User, error) {
	rows, err := s.Db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return types.User{}, err
	}

	for rows.Next() {
		return ScanUserRow(rows)
	}

	return types.User{}, fmt.Errorf("User not found")
}

func (s *PostgresStore) GetUserByEmail(email string) (types.User, error) {
	rows, err := s.Db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return types.User{}, err
	}

	for rows.Next() {
		return ScanUserRow(rows)
	}

	return types.User{}, fmt.Errorf("User not found")
}

func (s *PostgresStore) UpdateUser(user types.User) error {
	return nil
}

func (s *PostgresStore) DeleteUser(userId uuid.UUID) error {
	return nil
}

// Bookmark operations

func (s *PostgresStore) CreateBookmark(bookmark types.Bookmark) error {
	query := `
	insert into bookmarks
	(id, user_id, url, summary)
	values ($1, $2, $3, $4)
	`

	_, err := s.Db.Exec(
		query,
		bookmark.Id,
		bookmark.UserId,
		bookmark.Url,
		bookmark.Summary,
	)
	if err != nil {
		return err
	}

	log.Printf("Bookmark created\n%+v\n", utils.JSONMarshalIndent(bookmark))
	return nil
}

func (s *PostgresStore) GetBookmark(id uuid.UUID) (types.Bookmark, error) {
	rows, err := s.Db.Query("SELECT * FROM bookmarks WHERE id = $1", id)
	if err != nil {
		return types.Bookmark{}, err
	}

	for rows.Next() {
		return ScanBookmarkRow(rows)
	}

	return types.Bookmark{}, fmt.Errorf("Bookmark not found")
}

func (s *PostgresStore) GetBookmarksByUserId(userId uuid.UUID) ([]types.Bookmark, error) {
	rows, err := s.Db.Query("SELECT * FROM bookmarks WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}

	bookmarks := []types.Bookmark{}
	for rows.Next() {
		bookmark, err := ScanBookmarkRow(rows)
		if err != nil {
			return nil, err
		}

		bookmarks = append(bookmarks, bookmark)
	}

	return bookmarks, nil
}

func (s *PostgresStore) UpdateBookmark(bookmark types.Bookmark) error {
	query := `
	UPDATE bookmarks SET
	url = $2,
	summary = $3
	WHERE id = $1
	`

	_, err := s.Db.Exec(
		query,
		bookmark.Id,
		bookmark.Url,
		bookmark.Summary,
	)
	if err != nil {
		return err
	}

	log.Printf("Bookmark updated\n%+v\n", utils.JSONMarshalIndent(bookmark))
	return nil
}

func (s *PostgresStore) DeleteBookmark(id uuid.UUID) error {
	_, err := s.Db.Exec("DELETE FROM bookmarks WHERE id = $1", id)
	if err != nil {
		return err
	}

	log.Println("Bookmark deleted", id)
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

	if err := s.CreateBookmarksTable(); err != nil {
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

func (s *PostgresStore) CreateBookmarksTable() error {
	query := `CREATE TABLE IF NOT EXISTS bookmarks (
	id VARCHAR(36),
	user_id VARCHAR(36),
	url TEXT,
	summary TEXT,
	PRIMARY KEY(id),
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
	s.Db.Exec("DROP TABLE IF EXISTS bookmarks")
	s.Db.Exec("DROP TABLE IF EXISTS users")
}

// Helpers

func ScanSessionRow(rows *sql.Rows) (types.Session, error) {
	var session types.Session

	err := rows.Scan(
		&session.Token,
		&session.UserId,
		&session.ExpiresAt,
	)

	return session, err
}

func ScanUserRow(rows *sql.Rows) (types.User, error) {
	var user types.User

	err := rows.Scan(
		&user.Id,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	return user, err
}

func ScanBookmarkRow(rows *sql.Rows) (types.Bookmark, error) {
	var bookmark types.Bookmark

	err := rows.Scan(
		&bookmark.Id,
		&bookmark.UserId,
		&bookmark.Url,
		&bookmark.Summary,
	)

	return bookmark, err
}
