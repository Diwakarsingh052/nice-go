package main

import (
	"log"
	"os"
)

type logging struct {
	l *log.Logger

	// define a field that stores the connection to the logger struct from the log package
}

func main() {

	l := log.New(os.Stdout, "sales-app: ", log.Lshortfile)
	log := logging{l: l}
	l.Println("hello this is a custom log")
	log.l.Println("using custom type")

}
