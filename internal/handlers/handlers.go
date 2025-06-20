package handlers

import (
	"github.com/qwerun/habr-auth-go/pkg/postgres"
	"net/http"
)

type Server struct {
	explorer *postgres.Explorer
}

type registerRequest struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
	Nickname     string `json:"nickname"`
}

func NewMux(explorer *postgres.Explorer) (http.Handler, error) {
	server := &Server{explorer: explorer}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/register", server.register)
	mux.HandleFunc("/api/v1/verify-email", server.defaults)
	mux.HandleFunc("/api/v1/login", server.defaults)
	mux.HandleFunc("/api/v1/password-reset", server.defaults)
	return onlyPOST(mux), nil
}
