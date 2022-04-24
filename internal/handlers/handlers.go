package handlers

import (
	"github.com/zepyrshut/go-web-starter-gin/internal/config"
	"github.com/zepyrshut/go-web-starter-gin/internal/driver"
	"github.com/zepyrshut/go-web-starter-gin/internal/repository"
	"github.com/zepyrshut/go-web-starter-gin/internal/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	App *config.Application
	DB  repository.DBRepo
}

func NewRepo(a *config.Application, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
