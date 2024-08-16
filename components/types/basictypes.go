package types


type NewLoginRequest struct {
	Email  string `json:"email"`
	UserID string `json:"userID"`
}

type Response struct {
	UserID string `json:"user"`
	Email  string  `json:"email"`
	Profile bool `json:"profile"`
}

type NewUser struct{
	Email     string 
	UserID    string
	FirsName  string
	LastName  string
	Phone     string
	Interests []string
}