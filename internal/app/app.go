package app

import (
	"go-web/configs"
	"go-web/internal/database"
	"go-web/internal/handlers"
	"go-web/internal/repositories"
	"go-web/internal/services"
)

type App struct {
	DB *database.DB
	Config *configs.Config
	Address string
	// Repo, hdlr, svc
	UserRepo repositories.UserRepo
	SessionRepo repositories.SessionRepo
	UserSvc services.UserSvc
	UserHdl handlers.UserHdl
}

func InitApp() *App {
	cfg := configs.InitENV()
	db := database.InitDB(cfg)
	
	userRepo := repositories.NewUserRepo(db)
	sessionRepo := repositories.NewSessionRepo(db)
	userSvc := services.NewUserSvc(userRepo, sessionRepo)
	userHdl := handlers.NewUserHdl(userSvc)

	return &App{
		DB: db,
		Config: cfg,
		Address: configs.GetAppAddress(cfg),
		UserRepo: userRepo,
		SessionRepo: sessionRepo,
		UserSvc: userSvc,
		UserHdl: userHdl,
	}
}