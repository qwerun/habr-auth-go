package handlers

import (
	"github.com/qwerun/habr-auth-go/pkg/postgres"
	"net/http"
)

type Server struct {
	explorer *postgres.Explorer
}

func NewMux(explorer *postgres.Explorer) (http.Handler, error) {
	server := &Server{explorer: explorer}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/register", server.defaults)
	mux.HandleFunc("/api/v1/verify-email", server.defaults)
	mux.HandleFunc("/api/v1/login", server.defaults)
	mux.HandleFunc("/api/v1/password-reset", server.defaults)
	return mux, nil
}

func (s *Server) defaults(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Default"))
}
