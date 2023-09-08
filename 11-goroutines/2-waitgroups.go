package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//panic("work") // panic reveals go routines id
	wg := &sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		wg.Add(1)      // counter to add number of goroutines that we are starting or spinning up
		go work(i, wg) // each func call creates a goroutine
	}
	//tasks in the main
	// will wait until the counter resets to zero
	wg.Wait() // wait at the end of the main
}

func work(i int, wg *sync.WaitGroup) {
	defer wg.Done() // decrements the counter by one //
	wg1 := &sync.WaitGroup{}
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("some go routine")

	}()
	fmt.Println("I am doing some work", i)
	wg1.Wait()

}
