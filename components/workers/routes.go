package workers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harshgupta9473/zifty/components/middleware"
)

func (s *Server) Run() {
	routes := mux.NewRouter()

	routes.HandleFunc("/", middleware.WithJWTAuth(s.HandleHome))
	routes.HandleFunc("/editprofile",middleware.WithJWTAuth(s.HandleEditProfile))
	routes.HandleFunc("/signin", s.handleSignIN)
	routes.HandleFunc("/verify", s.handleVerification)

	log.Println("JSON API server is running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, routes)
}
