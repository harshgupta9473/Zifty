package workers

import (
	"encoding/json"
	"net/http"

	"github.com/harshgupta9473/zifty/components/db"
)

func NewServer(listenAddr string, store db.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

type Server struct {
	listenAddr string
	store      db.Storage
}

func WriteJSON(w http.ResponseWriter, status int, v any) {

	w.Header().Set(`Content-Type`, `application/json`)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)

}

type ServerError struct {
	Error string `json:"error"`
}
