package web

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// creating a custom type for storing the value in the context
type ctxKey int

const Key ctxKey = 1

// App struct is used to inject dependencies in the web package
type App struct {
	*chi.Mux
	*zap.Logger
}

// Values // we would put it in context so other layers can see traceId and when request started
type Values struct {
	TraceId    string
	Now        time.Time
	StatusCode int
}

// HandlerFunc is a custom type like http.HandlerFunc func in standard library
type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func (a *App) HandleFunc(method string, pattern string, handler HandlerFunc) {
	h := func(w http.ResponseWriter, r *http.Request) {

		v := &Values{
			TraceId: uuid.NewString(),
			Now:     time.Now(),
		}

		ctx := r.Context()
		// putting the values in the context
		ctx = context.WithValue(ctx, Key, v)

		//exec the handler
		err := handler(ctx, w, r)
		if err != nil {
			a.Logger.Error("error escaped from the middleware ", zap.Any("Error", err))
			//log.Println("error escaped from the middleware ", err)
			return
		}
	}

	a.Mux.MethodFunc(method, pattern, h)

}
