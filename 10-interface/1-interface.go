package main

import "fmt"

// Polymorphism means that a piece of code changes its behavior depending on the
//concrete data it’s operating on // Tom Kurtz, Basic inventor

// "Don’t design with interfaces, discover them". - Rob Pike
// Bigger the interface weaker the abstraction // Rob Pike

// add a method named as hello()
type reader interface {
	read(b []byte) (int, error)
}

type file struct {
	name string
}

func (f file) read(b []byte) (int, error) {
	fmt.Println("inside the file read")
	s := "hello go devs"
	copy(b, s)
	return len(b), nil
}

type json struct {
	data string
}

func (j json) read(b []byte) (int, error) {
	fmt.Println("inside the json read")
	s := `{name:"abc"}`
	copy(b, s)
	return len(b), nil
}

// fetch is a polymorphic func
// fetch() will accept any type of value which implements reader interface
func fetchData(r reader) {
	data := make([]byte, 50)
	r.read(data)
	fmt.Println(string(data))
	fmt.Println()
}

func main() {
	f := file{name: "abc.txt"} // concrete data
	fetchData(f)
	j := json{data: "any json"}
	fetchData(j)
}