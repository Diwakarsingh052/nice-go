package middleware

import "go.uber.org/zap"

type Mid struct {
	log *zap.Logger
}

func NewMid(log *zap.Logger) Mid {
	return Mid{
		log: log,
	}
}
