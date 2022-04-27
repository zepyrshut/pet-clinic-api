package dbrepo

import (
	"database/sql"

	"github.com/zepyrshut/pet-clinic/internal/config"
	"github.com/zepyrshut/pet-clinic/internal/repository"
)

type postgreDBRepo struct {
	App *config.Application
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, app *config.Application) repository.DBRepo {
	return &postgreDBRepo{
		App: app,
		DB:  conn,
	}
}
