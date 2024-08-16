package db

import "github.com/harshgupta9473/zifty/components/types"

type Storage interface {
	InsertIntoUserTable(types.NewUser) (*types.User, error)
	InsertIntoEmailVerificationTable(string, string, string) error
	UpdateProfile( types.NewUser)error


	VerifyTokenForEmailVerification(string, string,string) error

	GetUserByUserID(string) (*types.User, error)
	GetUserByEmail(string) (*types.User, error)
}
