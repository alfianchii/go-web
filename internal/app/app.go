package app

import (
	"go-boilerplate/configs"
	"go-boilerplate/internal/database"
)

type App struct {
	DB *database.DB
	// Repo, hdlr, svc
}

func InitApp() *App {
	cfg := configs.InitENV()
	db := database.InitDB(cfg)

	return &App{
		DB: db,
	}
}