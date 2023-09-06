package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		//always pass the value as argument to the go routine to have a local copy of it
		func(id int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second)
}
