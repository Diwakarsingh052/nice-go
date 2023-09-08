package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func main() {
	//http.HandleFunc()
	HandleFunc("/home", homeH)
	http.ListenAndServe(":8080", nil)
}

func HandleFunc(pattern string, handler HandlerFunc) {
	// h var would have the same signature as what http. HandleFunc accepts // inside it, we are calling our custom handler
	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := handler(ctx, w, r)
		if err != nil {
			log.Println("some error happened")
			return
		}

	} // the function is constructed not called

	http.HandleFunc(pattern, h) // http package would call the h function internally,
	// h function contains our custom logic
}

func homeH(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	log.Println("custom handler func exec")
	fmt.Fprintln(w, "hello this is a home page")
	return nil
}
