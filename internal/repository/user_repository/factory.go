package user_repository

import "github.com/qwerun/habr-auth-go/pkg/postgres"

type Repository struct {
	explorer *postgres.Explorer
}

func New(explorer *postgres.Explorer) *Repository {
	return &Repository{explorer: explorer}
}
