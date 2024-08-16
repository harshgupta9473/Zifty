package db


func (s *PostgresStore) CreateUsersTable() error {
	query := `create table if not exists users(
	id integer generated always as identity primary key,
	email varchar(255) unique not null,
	userid varchar(100) unique not null,
	firstname varchar(100) not null,
	lastname varchar(100),
	phoneNumber varchar(20) not null,
	interests jsonb,
	verified boolean default false
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateEmalVerificationTable() error {
	query := `create table if not exists emailverification(
	emailid varchar(100) not null,
	userID varchar(100)  not null,
	token varchar(100) not null,
	expires_at timestamp not null
	)`
	_, err := s.db.Exec(query)
	return err
}

