package middleware

import (
	"go.uber.org/zap"
	"service-app/auth"
)

type Mid struct {
	log *zap.Logger
	a   *auth.Auth
}

func NewMid(log *zap.Logger, a *auth.Auth) Mid {
	return Mid{
		log: log,
		a:   a,
	}
}
