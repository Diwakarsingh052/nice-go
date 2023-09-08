package main

import (
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

		next(w, r.WithContext(ctx))

	}
}
//LoggingMid // figure out the signature
func LoggingMid() {
	//return correct value
	return {
		// check context is set or not // get the value out of the context and store in a variable
		ctx :=  r.Context()

		if !ok {
			reqID = "unknown"
		}
		log.Printf("%s : started   : %s %s ",
			reqID,
			r.Method, r.URL.Path)

		next(w, r) // calling the next thing in the chain
		log.Printf("%s : completed : %s %s ",
			reqID,
			r.Method, r.URL.Path,

	}
}
