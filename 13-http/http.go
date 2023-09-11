package main

import (
	"net/http"
)

func main() {
	//Home is handlerFunc
	// http.HandleFunc register a route and accept handlerFunc which will handle the request for /home endpoint
	http.HandleFunc("/home", Home)
	//nil means we are using a default handler
	panic(http.ListenAndServe(":8080", nil))
}

// Home
// http.ResponseWriter is used to write response back to the client ,
// http.Request is used to get any request-specific details like json, req body , or anything related to request data

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("hello this is our first home page"))
}
