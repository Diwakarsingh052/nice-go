package middleware

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"runtime/debug"
	"service-app/web"
)

func (m *Mid) Panic(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
		// type assertion and making sure values struct with the trace id is present
		v, ok := ctx.Value(web.Key).(*web.Values)
		if !ok {
			return fmt.Errorf("web.Values missing from the context")
		}
		defer func() {
			r := recover()
			if r != nil { // panic happened
				s := fmt.Sprintf("PANIC :%v", r)
				err = errors.New(s)
				m.log.Info("PANIC", zap.Any("Trace Id", v.TraceId), zap.String("StackTrace", string(debug.Stack())))
			}
		}()
		return next(ctx, w, r)
	}

}
