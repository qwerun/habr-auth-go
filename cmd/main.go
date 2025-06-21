package main

import (
	"github.com/qwerun/habr-auth-go/internal/handlers"
	"github.com/qwerun/habr-auth-go/internal/repository/user_repository"
	"github.com/qwerun/habr-auth-go/pkg/postgres"
	"log"
	"net/http"
)

func main() {
	db, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	explorer := postgres.NewExplorer(db)
	userRepo := user_repository.New(explorer)
	handler, err := handlers.NewMux(userRepo)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8081", handler)

}
