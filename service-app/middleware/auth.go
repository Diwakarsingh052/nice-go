package middleware

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"service-app/auth"
	"service-app/web"
	"strings"
)

func (m *Mid) Authenticate(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		// type assertion and making sure values struct with the trace id is present
		v, ok := ctx.Value(web.Key).(*web.Values)
		if !ok {
			return fmt.Errorf("web.Values missing from the context")
		}

		authHeader := r.Header.Get("Authorization")
		//token format :- Bearer <token>
		parts := strings.Split(authHeader, " ") // parts would be slice

		//making sure if the bearer was passed with the token
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			err := errors.New("expected authorization header format: Bearer <token>")
			return web.NewRequestError(err, http.StatusUnauthorized)

		}
		//parts[1] of contains the token
		claims, err := m.a.ValidateToken(parts[1])
		if err != nil {
			//this is for internal logs, end user would not see it
			m.log.Error(err.Error(), zap.Any("Trace Id", v.TraceId))

			webErr := errors.New(http.StatusText(http.StatusUnauthorized))
			//error message sent to the external user
			return web.NewRequestError(webErr, http.StatusUnauthorized)
		}

		// put claims in the context
		ctx = context.WithValue(ctx, auth.Key, claims)

		return next(ctx, w, r)
	}
}

// Authorize checks if user has enough permission to call the endpoint
func (m *Mid) Authorize(next web.HandlerFunc, requiredRoles ...string) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

		// type assertion and making sure values struct with the trace id is present
		v, ok := ctx.Value(web.Key).(*web.Values)
		if !ok {
			return fmt.Errorf("web.Values missing from the context")
		}

		//fetching claims data from the context
		claims, ok := ctx.Value(auth.Key).(auth.Claims)

		if !ok {
			err := errors.New("claims missing from context: Authorize called without/before Authenticate")
			m.log.Error(err.Error(), zap.Any("Trace Id", v.TraceId))

			webErr := errors.New(http.StatusText(http.StatusUnauthorized))
			//error message sent to the external user
			return web.NewRequestError(webErr, http.StatusUnauthorized)

		}

		ok = claims.HasRoles(requiredRoles...) // pass the individual values of the requiredRoles to the hasRoles
		// variadic
		// param

		if !ok {
			webErr := errors.New(http.StatusText(http.StatusUnauthorized))
			//error message sent to the external user
			return web.NewRequestError(webErr, http.StatusUnauthorized)
		}

		return next(ctx, w, r)
	}
}
