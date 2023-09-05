package main

import (
	"fmt"
	"log"
)

type user struct {
	name  string
	email string
}

func (u user) Write(p []byte) (n int, err error) {
	fmt.Printf("sending a notification to %s %s %s", u.name, u.email, string(p))
	return len(p), nil
}

func main() {
	var u user
	// we need to implement the writer interface to pass user object to log.New function
	l := log.New(u, "sales-app: ", log.Lshortfile) // remove the compile time error
	l.Println("hello")
}
