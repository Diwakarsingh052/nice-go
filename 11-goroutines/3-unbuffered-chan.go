package main

import (
	"fmt"
	"sync"
)

// https://go.dev/ref/spec#Send_statements
// A send on an unbuffered channel can proceed if a receiver is ready. send will block until there is no recv
func main() {
	//unbuffered //buffered
	wg := &sync.WaitGroup{}
	c := make(chan int) //unbuffered
	wg.Add(2)
	go func() {
		defer wg.Done()
		c <- 10 // send
		//print sending is finished
	}()

	//x := <-c // receive
	go func() {
		defer wg.Done()
		//sleep here for 3 seconds
		fmt.Println(<-c) // blocking call for the current goroutine
	}()
	wg.Wait()
}
