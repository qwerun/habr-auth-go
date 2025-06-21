package main

import (
	"github.com/qwerun/habr-auth-go/internal/handlers"
	"github.com/qwerun/habr-auth-go/internal/repository/user_repository"
	"github.com/qwerun/habr-auth-go/pkg/postgres"
	"github.com/qwerun/habr-auth-go/pkg/redis"
	"log"
	"net/http"
)

func main() {
	db, err := postgres.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	rdb, err := redis.NewRedisDB()
	if err != nil {
		log.Fatal(err)
	}
	rExplorer := redis.NewRedisExplorer(rdb)
	explorer := postgres.NewExplorer(db)
	userRepo := user_repository.New(explorer, rExplorer)
	handler, err := handlers.NewMux(userRepo)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8081", handler)

}
