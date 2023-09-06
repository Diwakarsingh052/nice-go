package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1) //  // counter to add number of goroutines that we are starting or spinning up
	for i := 1; i <= 10; i++ {
		work(i, wg) // make this func call concurrent
	}
	//tasks in main
	// it will wait until counter resets to zero
	wg.Wait() // wait at the end of the main
}

func work(i int, wg *sync.WaitGroup) {
	fmt.Println("I am doing some work", i)
	wg.Done() // decrements the counter by one //
}
