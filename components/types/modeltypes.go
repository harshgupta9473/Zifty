package types

import "time"

type User struct {
	ID        uint64
	Email     string
	UserID    string
	FirsName  string
	LastName  string
	Phone     string
	Interests []string
	Verified  bool
}

type Emailerification struct {
	Email     string
	UserID    string
	Token     string
	ExpiresAt time.Time
}
