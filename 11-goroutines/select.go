package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = &sync.WaitGroup{}
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c3 := make(chan string, 1)
	done := make(chan struct{})

	//wgWorker keep track of if the go routine work is finished or not and we wil close the channel when work is done
	var wgWorker = &sync.WaitGroup{}

	wgWorker.Add(3)
	go func() {
		defer wgWorker.Done()
		time.Sleep(time.Second * 3)
		c1 <- "one" // send
	}()
	go func() {
		defer wgWorker.Done()
		time.Sleep(time.Second)
		c2 <- "two" // send
	}()

	go func() {
		defer wgWorker.Done()
		c3 <- "three" // send
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wgWorker.Wait()
		close(done) // close the channel when goroutines are finished sending
	}()

	//fmt.Println(<-c1)
	//fmt.Println(<-c2)
	//fmt.Println(<-c3)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			// whichever case is not blocking exec that first
			//whichever case is ready first exec that.
			//case chan recv , send , default
			case x := <-c1:
				fmt.Println(x)
			case y := <-c2:
				fmt.Println(y)
			case z := <-c3:
				fmt.Println(z)
			case <-done: // this case will exec when channel is closed
				fmt.Println("it is closed,work is finished")
				return

			}
		}
	}()

	fmt.Println("end of the main")
	wg.Wait()

}
