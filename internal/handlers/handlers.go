package handlers

import (
	"encoding/json"
	"github.com/qwerun/habr-auth-go/pkg/postgres"
	"net/http"
	"strings"
)

type Server struct {
	explorer *postgres.Explorer
}

type User struct {
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

func onlyPOST(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	var req User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}
	if err = req.IsValid(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := req
	if err = writeJSON(w, response, http.StatusOK); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) defaults(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Default"))
}

func writeJSON(w http.ResponseWriter, v any, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
