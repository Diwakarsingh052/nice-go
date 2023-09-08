package handlers

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"service-app/web"
)

func API(log *zap.Logger) http.Handler {
	app := web.App{
		Mux:    chi.NewRouter(),
		Logger: log,
	}

	app.HandleFunc(http.MethodGet, "/check", check)
	return app

}
