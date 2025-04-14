package app

import (
	"go-web/configs"
	"go-web/internal/database"
	"go-web/internal/repositories"
)

type App struct {
	DB *database.DB
	Config *configs.Config
	Address string
	// Repo, hdlr, svc
	UserRepo repositories.UserRepo
}

func InitApp() *App {
	cfg := configs.InitENV()
	db := database.InitDB(cfg)
	
	userRepo := repositories.NewUserRepo(db)

	return &App{
		DB: db,
		Config: cfg,
		Address: configs.GetAppAddress(cfg),
		UserRepo: userRepo,
	}
}