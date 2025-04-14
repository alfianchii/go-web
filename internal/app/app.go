package app

import (
	"go-web/configs"
	"go-web/internal/database"
)

type App struct {
	DB *database.DB
	Config *configs.Config
	Address string
	// Repo, hdlr, svc
}

func InitApp() *App {
	cfg := configs.InitENV()
	db := database.InitDB(cfg)

	return &App{
		DB: db,
		Config: cfg,
		Address: configs.GetAppAddress(cfg),
	}
}