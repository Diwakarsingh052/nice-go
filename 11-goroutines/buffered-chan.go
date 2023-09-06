package main

import (
	"fmt"
	"sync"
)

// A send on a buffered channel can proceed if there is room in the buffer.
func main() {
	wg := &sync.WaitGroup{}
	c := make(chan int, 2) //buffered
	wg.Add(1)
	go func() {
		defer wg.Done()
		c <- 10 // send
		c <- 10
		c <- 10
		fmt.Println("sending finished")
	}()

	//x := <-c // receive
	go func() {
		defer wg.Done()
		//chan recv make the room in the buffer
		fmt.Println(<-c) // blocking call for the current goroutine
	}()
	wg.Wait()
}
