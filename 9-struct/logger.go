package main

import (
	"log"
)

type logging struct {
	name string
	l    *log.Logger
	// define a field that stores the connection to the logger struct from the log package
}

func NewLogging(name string, l *log.Logger) logging {
	return logging{
		name: name,
		l:    l,
	}
}

func main() {
	log.New()
	//l := log.New(os.Stdout, "sales-app: ", log.Lshortfile)
	//log := logging{l: l}
	//l.Println("hello this is a custom log")
	//log.l.Println("using custom type")

}
