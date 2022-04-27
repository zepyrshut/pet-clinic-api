package handlers

import (
	"github.com/zepyrshut/pet-clinic/internal/config"
	"github.com/zepyrshut/pet-clinic/internal/database"
	"github.com/zepyrshut/pet-clinic/internal/repository"
	"github.com/zepyrshut/pet-clinic/internal/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	App *config.Application
	DB  repository.DBRepo
}

func NewRepo(a *config.Application, db *database.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
