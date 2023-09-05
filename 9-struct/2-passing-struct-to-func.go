package main

import (
	"fmt"
	"net/http"
)

type config struct {
	hostName string
	port     int
	method   string
}

func main() {
	var c config
	runApp(10, c)

}

// runApp defines config param which help us to achieve default values in function
func runApp(abc int, c config) {
	if c.port == 0 {
		c.port = 8080
	}
	if c.hostName == "" {
		c.hostName = "localhost"
	}
	if c.method == "" {
		c.method = http.MethodGet
	}

	fmt.Printf("%+v\n", c)
}
