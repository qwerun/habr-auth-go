package handlers

import (
	"encoding/json"
	"github.com/qwerun/habr-auth-go/internal/auth"
	"github.com/qwerun/habr-auth-go/internal/models"
	"log"
	"net/http"
)

func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}
	if err = req.IsValid(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hashed, err := auth.HashPassword(req.PasswordHash)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	req.PasswordHash = hashed

	user := models.NewUser(req.Email, req.PasswordHash, req.Nickname)
	var id string
	if id, err = s.explorer.Create(user); err != nil {
		log.Print(err.Error())
		http.Error(w, "Failed register", http.StatusInternalServerError)
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
