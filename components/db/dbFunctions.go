package db

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/harshgupta9473/zifty/components/types"
)

func (s *PostgresStore) InsertIntoUserTable(user types.NewUser) (*types.User, error) {
	query := `insert into users
	(email,userID,firstname,lastname,phoneNumber,interests,verified)
	values($1,$2,$3,$4,$5,$6,$7)`
	interestsJSON, err := json.Marshal(user.Interests)
	if err != nil {
		return nil, err
	}
	_, err = s.db.Exec(query, user.Email, user.UserID, user.FirsName, user.LastName, user.Phone, interestsJSON, true)
	if err != nil {
		return nil, err
	}
	userPro, err := s.GetUserByEmail(user.UserID)
	if err!=nil{
		return nil,err
	}
	return userPro,nil

}

func (s *PostgresStore) InsertIntoEmailVerificationTable(emailID string, userID string, token string) error {
	expiration := time.Now().Add(5 * time.Minute)
	query := `insert into emailverification(emailid,userid,token,expires_at)
	values($1,$2,$3,$4)`
	_, err := s.db.Exec(query, emailID, userID, token, expiration)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) VerifyTokenForEmailVerification(email, userID, token string) error {
	var expiretime time.Time
	fmt.Println(token)
	
	err := s.db.QueryRow("select expires_at from emailverification where emailid=$1 and userid=$2 and token=$3", email,userID,token).Scan(&expiretime)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if time.Now().After(expiretime) {
		return fmt.Errorf("token expired")
	}
	return nil
}


func (s *PostgresStore)UpdateProfile(user types.NewUser)error{
	query:=`update users set userid=$1,firstname=$2,lastname=$3,phoneNumber=$4,interests=$5,verified=$6 where email=$7`

	interestsJSON, err := json.Marshal(user.Interests)
	if err != nil {
		return err
	}
	_,err=s.db.Exec(query,user.FirsName,user.LastName,user.Phone,interestsJSON,true,user.Email)
	if err!=nil{
		return err
	}
	return nil
}





func (s *PostgresStore) GetUserByUserID(userID string) (*types.User, error) {
	query := `select id, email,userID,firstname,lastname,phoneNumber,interests,verified from users where userID=$1 `
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var user = new(types.User)
	var interestsJSON []byte
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.UserID, &user.FirsName, &user.LastName, &user.Phone, &interestsJSON, &user.Verified)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("user not found")
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	err = json.Unmarshal(interestsJSON, &user.Interests)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *PostgresStore) GetUserByEmail(email string) (*types.User, error) {
	query := `select id, email,userID,firstname,lastname,phoneNumber,interests,verified from users where email=$1 `
	rows, err := s.db.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var user = new(types.User)
	var interestsJSON []byte
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.UserID, &user.FirsName, &user.LastName, &user.Phone, &interestsJSON, &user.Verified)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("user not found")
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	err = json.Unmarshal(interestsJSON, &user.Interests)
	if err != nil {
		return nil, err
	}
	return user, nil

}
