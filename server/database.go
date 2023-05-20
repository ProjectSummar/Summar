package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB interface {
	CreateUser(*User) error
	// DeleteUser(string) error
	// UpdateUser(*User) error
	GetUserByEmail(string) (*User, error)
	// GetUserById(string) (*User, error)
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
	return nil, nil
}

// Initialisation

func (db *PostgresDB) Init() error {
	if err := db.CreateUserTable(); err != nil {
		return err
	}

	return nil
}

func (db *PostgresDB) CreateUserTable() error {
	query := `create table if not exists users (
	id varchar(36) primary key,
	email varchar(50),
	password_hash varchar(60),
	created_at timestamp
	)`

	_, err := db.Db.Exec(query)
	return err
}