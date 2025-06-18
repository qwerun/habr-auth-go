package main

import (
	"github.com/qwerun/habr-auth-go/api"
	"github.com/qwerun/habr-auth-go/internal/db/postgres"
	"log"
	"net/http"
)

func main() {
	db, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	explorer := postgres.NewExplorer(db)
	handler, err := api.NewMux(explorer)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8081", handler)

}
