package main

import "fmt"

func main() {

	show("hello", 10, 20, 30, 40, 50)
}

func show(s string, i ...int) {
	fmt.Printf("%T\n", i)
	fmt.Println(s, i)

}
