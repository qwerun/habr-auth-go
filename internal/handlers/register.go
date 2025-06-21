package handlers

import (
	"encoding/json"
	"errors"
	"github.com/qwerun/habr-auth-go/internal/auth"
	"github.com/qwerun/habr-auth-go/internal/models"
	"github.com/qwerun/habr-auth-go/internal/repository/user_repository"
	"log"
	"net/http"
)

func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Bad JSON: %v", err)
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}
	if err = req.IsValid(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hashed, err := auth.HashPassword(req.PasswordHash)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	req.PasswordHash = hashed

	user := models.NewUser(req.Email, req.PasswordHash, req.Nickname)
	var id string
	if id, err = s.explorer.Create(user); err != nil {
		switch {
		case errors.Is(err, user_repository.ErrEmailAlreadyExists):
			http.Error(w, err.Error(), http.StatusConflict)
		case errors.Is(err, user_repository.ErrNicknameAlreadyExists):
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	req.Email = id

	response := req

	if err = writeJSON(w, response, http.StatusOK); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) defaults(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Default"))
}
