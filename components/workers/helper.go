package workers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateToken() (string, error) {
	token := make([]byte, 16)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}

func extractFromJWT(token *jwt.Token)(string,string,error){
	claims:=token.Claims.(jwt.MapClaims)
	userID,ok:=claims["userID"].(string)
	if !ok{
		return "","" ,fmt.Errorf("error")
	}
	email,ok:=claims["user"].(string)
	if !ok{
		return "","" ,fmt.Errorf("error")
	}
	return email,userID,nil
}

