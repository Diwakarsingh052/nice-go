package main

import (
	"fmt"
	"time"
)

func main() {

	go hello()  // spins up a go routine
	go func() { // spinning up an anonymous goroutine
		fmt.Println("task 1")
		fmt.Println("taks 2")
		fmt.Println("task 3 ")
	}() // () this means we are calling the function
	fmt.Println("doing tasks specific to main")
	fmt.Println("end of the main")
	time.Sleep(1 * time.Second)
}

func hello() {
	fmt.Println("hello from the hello func")
}
