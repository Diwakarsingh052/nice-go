package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//note:- this below pattern not recommended
	wg := sync.WaitGroup{}
	ch := make(chan int, 4)

	wg.Add(1)
	go func() {
		defer wg.Done()
		//range will iterate over the channel // recv all the values from the channel until sender is sending
		for data := range ch {

			fmt.Println(data)
		}
	}()
	time.Sleep(time.Second)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch) // close the channel for further send //for range recv will continue normally, and it will empty the
	// channel and quit
	//ch <- 20

	wg.Wait()
}
