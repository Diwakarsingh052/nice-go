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
	wg     sync.WaitGroup
}

func (c *cabs) getCabDriver() {
	defer c.wg.Done()
	// read lock:when a goroutine is reading then no other can be writing to the shared resource//
	// there could be unlimited number of reads
	c.rw.RLock()
	defer c.rw.RUnlock()
	fmt.Println("driver", c.driver)

}
func (c *cabs) bookCab(name string) {
	defer c.wg.Done()
	fmt.Println("welcome to the website", name)
	fmt.Println("some offers for you", name)

	// write lock // no one can read if a goroutine is writing //only one goroutine can enter to write.
	c.rw.Lock()
	// when a goroutine acquires a lock then another go routine can't access the critical section
	//until the lock is not released
	defer c.rw.Unlock()

	//critical section
	if c.driver >= 1 {
		fmt.Println("car is available for", name)
		time.Sleep(3 * time.Second)
		fmt.Println("booking confirmed", name)
		c.driver--
	} else {
		fmt.Println("car is not available for", name)
	}
	//critical section ends
}

func main() {

	c := cabs{
		driver: 5,
	}

	for i := 1; i <= 5; i++ {
		c.wg.Add(1)
		go c.getCabDriver()
	}
	for i := 1; i <= 5; i++ {
		c.wg.Add(1)
		go c.bookCab("data " + strconv.Itoa(i))

	}
	for i := 1; i <= 5; i++ {
		c.wg.Add(1)
		go c.getCabDriver()
	}

	c.wg.Wait()
}
