package middleware

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"service-app/web"
	"time"
)

func (m *Mid) Logger(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		// type assertion and making sure values struct with the trace id is present
		v, ok := ctx.Value(web.Key).(*web.Values)
		if !ok {
			return fmt.Errorf("web.Values missing from the context")
		}

		m.log.Info("started", zap.Any("Trace Id", v.TraceId),
			zap.String("Method", r.Method), zap.Any("URL Path", r.URL.Path))

		err := next(ctx, w, r) // executing the next handlerFunc or the middleware in the chain

		m.log.Info("completed",
			zap.Any("Trace Id", v.TraceId),
			zap.String("Method", r.Method), zap.Any("URL Path", r.URL.Path),
			zap.Int("Status Code", v.StatusCode), zap.Duration("duration", time.Since(v.Now)))
		if err != nil {
			return err
		}
		return nil
	}
}
