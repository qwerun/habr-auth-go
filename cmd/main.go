package main

import (
	"github.com/qwerun/habr-auth-go/internal/handlers"
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
	handler, err := handlers.NewMux(explorer)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8081", handler)

}
