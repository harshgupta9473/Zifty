package workers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/harshgupta9473/zifty/components/middleware"
	"github.com/harshgupta9473/zifty/components/types"
)

func (s *Server) HandleHome(w http.ResponseWriter, r *http.Request, token *jwt.Token) {
	userID, email, err := extractFromJWT(token)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ServerError{Error: "server error"})
		return
	}
	user, err := s.store.GetUserByEmail(email)
	if err != nil {
		WriteJSON(w, http.StatusOK, types.Response{UserID: userID,Email: email,Profile: false})
		return
	}
	WriteJSON(w, http.StatusOK, user)
}

func (s *Server) HandleEditProfile(w http.ResponseWriter, r *http.Request, token *jwt.Token) {
	var newuser types.NewUser
	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ServerError{Error: err.Error()})
		return
	}
	_,err=s.store.GetUserByEmail(newuser.Email)
	if err==nil{
		err=s.store.UpdateProfile(newuser)
		if err!=nil{
			WriteJSON(w, http.StatusInternalServerError, ServerError{Error: err.Error()})
		return
		}
		user,err:=s.store.GetUserByEmail(newuser.Email)
		if err!=nil{
			WriteJSON(w, http.StatusInternalServerError, ServerError{Error: err.Error()})
		}
		WriteJSON(w,http.StatusOK,user)
		return
	}
	user, err := s.store.InsertIntoUserTable(newuser)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ServerError{Error: err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, user)
}

func (s *Server) handleSignIN(w http.ResponseWriter, r *http.Request) {
	var userReq types.NewLoginRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		log.Println(fmt.Errorf("not of correct format"))
		return
	}

	err = s.UserVerification(userReq)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ServerError{Error: err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, "Email Sent for the verification")

}

func (s *Server) handleVerification(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	email := r.URL.Query().Get("user")
	userID := r.URL.Query().Get("userid")
	if token == "" {
		WriteJSON(w, http.StatusBadRequest, "token is required")
		return
	}
	err := s.store.VerifyTokenForEmailVerification(email, userID, token)
	if err != nil {
		WriteJSON(w, http.StatusForbidden, err)
		return
	}

	jwtToken, err := middleware.CreateJWT(email, userID)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ServerError{Error: err.Error()})
		return
	}
	_, err = s.store.GetUserByEmail(email)
	var profile bool = true
	if err != nil {
		profile = false
	}

	s.setCookie(w, jwtToken)
	res := types.Response{
		UserID:  userID,
		Email:   email,
		Profile: profile,
	}

	WriteJSON(w, http.StatusOK, res)
}

func (s *Server) setCookie(w http.ResponseWriter, jwtToken string) {
	cookie := http.Cookie{
		Name:     "authToken",
		Value:    jwtToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		// Secure:   true,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)

}

func (s *Server) UserVerification(userReq types.NewLoginRequest) error {
	token, err := GenerateToken()
	if err != nil {
		return err
	}
	err = s.store.InsertIntoEmailVerificationTable(userReq.Email, userReq.UserID, token)
	if err != nil {
		return err
	}
	err = middleware.SendVerificationEmail(userReq.Email, userReq.UserID, token)
	if err != nil {
		return err
	}
	return nil
}
