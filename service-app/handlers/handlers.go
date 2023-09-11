package handlers

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"service-app/auth"
	"service-app/core/inventory"
	"service-app/core/user"
	"service-app/middleware"
	"service-app/web"
)

type handler struct {
	log *zap.Logger
	us  *user.Service
	inv *inventory.Service
	*auth.Auth
}

// API register the routes
func API(log *zap.Logger, a *auth.Auth, us *user.Service, inv *inventory.Service) http.Handler {
	//adding values to the app struct //
	app := web.App{
		Mux:    chi.NewRouter(),
		Logger: log,
	}
	m := middleware.NewMid(log, a)

	middlewares := []func(web.HandlerFunc) web.HandlerFunc{
		m.Logger,
		m.Error,
		m.Panic,
	}
	h := handler{
		log:  log,
		us:   us,
		inv:  inv,
		Auth: a,
	}

	//HandleFunc is the custom implementation // it is defined over the app struct
	//app.HandleFunc(http.MethodGet, "/check", m.Logger(m.Error(m.Panic(check))))
	//app.HandleFunc(http.MethodGet, "/check", m.Logger(m.Error(m.Panic(m.Authenticate(m.Authorize(check,
	//	auth.RoleAdmin, auth.RoleUser))))))

	app.HandleFunc(http.MethodGet, "/check", ChainMiddleware(m.AuthenticateCookie(m.Authorize(check, auth.RoleAdmin,
		auth.RoleUser)), middlewares))
	app.HandleFunc(http.MethodPost, "/signup", m.Logger(m.Error(m.Panic(h.Signup))))
	app.HandleFunc(http.MethodPost, "/login", m.Logger(m.Error(m.Panic(h.Login))))
	app.HandleFunc(http.MethodPost, "/add", ChainMiddleware(m.AuthenticateCookie(m.Authorize(h.AddInventory, auth.RoleAdmin,
		auth.RoleUser)), middlewares))
	//we can return the app struct as it implements the http.Handler interface
	return app

}

func ChainMiddleware(handler web.HandlerFunc, middlewares []func(handlerFunc web.HandlerFunc) web.HandlerFunc) web.
	HandlerFunc {

	// Apply middlewares in reverse so they execute in the order they're defined
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	/*
		handler = LogMiddleware(AuthMiddleware(handler))
		(Step 1)
		[AuthMiddleware] ---> [handler]
		Result: [AuthMiddleware[handler]]

		(Step 2)
		[LogMiddleware] ---> [AuthMiddleware[handler]]
		Result: [LogMiddleware[AuthMiddleware[handler]]]
	*/

	return handler
}
