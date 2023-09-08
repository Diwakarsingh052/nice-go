package main

import (
	"fmt"
	"log"
	"net/http"
)

// req -> handlerFunc() -> CreateUser // normal flow
// req -> mid -> handlerFunc() -> CreateUser

func main() {

	//Mid func return type should evaluate to func (w http.ResponseWriter,
	//r *http.Request) because it is called not passed as value
	http.HandleFunc("/home", Mid(home))
	panic(http.ListenAndServe(":8080", nil))

}

func home(w http.ResponseWriter, r *http.Request) {

	log.Println("In home Page handler")
	fmt.Fprintln(w, "this is my home page")

}

// Mid accepts a handler function
func Mid(next http.HandlerFunc) http.HandlerFunc {
	//type HandlerFunc func(ResponseWriter, *Request)
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("middleware is called")
		log.Println("doing middleware specific things")
		next(w, r)
		log.Println("middleware ended")
	}
}
