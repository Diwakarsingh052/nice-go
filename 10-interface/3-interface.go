package main

import (
	"fmt"
)

// library code that is already there
type writer interface {
	write()
}
type network struct {
	writer
}

func (n network) write() {
	fmt.Println("writing using the network struct")

}

// library ends

type abc struct {
	network // embedding the struct which already impl the interface // abc struct now also impl the interface
	// by inner type promotion
}

func handleNetworkRequest(a abc) writer {
	return a
}

func main() {
	var a abc
	a.write()
	var n network

}
