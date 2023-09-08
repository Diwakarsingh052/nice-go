package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// req -> handlerFunc() -> CreateUser // normal flow
// req -> mid -> handlerFunc() -> CreateUser
// go get github.com/google/uuid
// go mod tidy // remove any unused dependency, it will download dependencies listed in go.mod files,
// update the go.mod file for correct usage of module

type key int

const RequestIDKey key = 123

func main() {

	//Mid func return type should evaluate to func (w http.ResponseWriter,
	//r *http.Request) because it is called not passed as value
	http.HandleFunc("/home", RequestMid(LoggingMid(homePage)))
	panic(http.ListenAndServe(":8080", nil))

}

func homePage(w http.ResponseWriter, r *http.Request) {

	log.Println("In home Page handler")
	fmt.Fprintln(w, "this is my home page")

}

// RequestMid accepts a handler function
func RequestMid(next http.HandlerFunc) http.HandlerFunc {
	//type HandlerFunc func(ResponseWriter, *Request)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceId := uuid.NewString()
		// put the traceId in the context
		ctx = context.WithValue(ctx, RequestIDKey, traceId)
		next(w, r.WithContext(ctx))

	}
}
func LoggingMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//do middleware specific stuff first and when it is over then go to the actual handler func to exec it
		// fetching requestId value from the context and making sure the type store in it is of the string type
		reqID, ok := r.Context().Value(RequestIDKey).(string)
		if !ok {
			reqID = "unknown"
		}
		log.Printf("%s : started   : %s %s ",
			reqID,
			r.Method, r.URL.Path)
		if r.Method != http.MethodGet {
			http.Error(w, "method must be get", http.StatusInternalServerError)
			return
		}

		next(w, r) // executing the next handlerFunc or the middleware in the chain

		log.Printf("%s : completed : %s %s ",
			reqID,
			r.Method, r.URL.Path,
		)

	}
}
