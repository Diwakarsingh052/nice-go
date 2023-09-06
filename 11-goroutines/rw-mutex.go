package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type cabs struct {
	driver int
	rw     sync.RWMutex
}

func (c *cabs) getCabDriver(wg *sync.WaitGroup) {
	//Rlock // Runlock
	fmt.Println("driver", c.driver)
}

func (c *cabs) bookCab(name string, wg *sync.WaitGroup) {
	fmt.Println("welcome to the website", name)
	fmt.Println("some offers for you", name)
	//lock // unlock
	if c.driver >= 1 {
		fmt.Println("car is available for", name)
		time.Sleep(3 * time.Second)
		fmt.Println("booking confirmed", name)
		c.driver--
	} else {
		fmt.Println("car is not available for", name)
	}
}
func main() {
	wg := &sync.WaitGroup{}
	c := cabs{
		driver: 5,
	}
	for i := 1; i <= 15; i++ {
		go c.getCabDriver(wg)
	}

	for i := 1; i <= 15; i++ {
		go c.bookCab("user "+strconv.Itoa(i), wg)
	}

	for i := 1; i <= 15; i++ {
		go c.getCabDriver(wg)
	}

}
