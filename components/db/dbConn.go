package db

import (
	"database/sql"
	"os"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	connstr := os.Getenv("dockerConnStr")
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	err := s.CreateUsersTable()
	if err != nil {
		return err
	}
	err = s.CreateEmalVerificationTable()
	if err != nil {
		return err
	}
	return nil
}
