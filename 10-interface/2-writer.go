package main

import "log"

type user struct {
	name  string
	email string
}

func main() {
	var u user
	l := log.New(u, "sales-app: ", log.Lshortfile) // remove the compile time error
	l.Println("hello")
}
