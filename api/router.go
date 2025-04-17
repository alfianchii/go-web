package api

import (
	"go-web/internal/app"
	"go-web/internal/middlewares"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func InitRouter(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		// Routes
		r.Post("/login", app.UserHdl.Login)

		r.Group(func(r chi.Router) {
			r.Use(middlewares.AuthMiddleware("admin", app.UserSvc, app.SessionRepo))
			r.Get("/dashboard", app.DashboardHdl.DashboardData)
		})
	})

	return r
}