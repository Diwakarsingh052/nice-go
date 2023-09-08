package handlers

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"service-app/middleware"
	"service-app/web"
)

// API register the routes
func API(log *zap.Logger) http.Handler {
	//adding values to the app struct //
	app := web.App{
		Mux:    chi.NewRouter(),
		Logger: log,
	}
	m := middleware.NewMid(log)

	//HandleFunc is the custom implementation // it is defined over the app struct
	app.HandleFunc(http.MethodGet, "/check", m.Logger(m.Error(check)))

	//we can return the app struct as it implements the http.Handler interface
	return app

}
