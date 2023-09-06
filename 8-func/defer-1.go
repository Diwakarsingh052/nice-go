package main

import "fmt"

func main() {
	//defer maintains a stack. [2,1]
	defer func() {
		defer fmt.Println(1) // when your function is returning defer statements will exec
		defer fmt.Println(2)
	}()
	panic("main ") // defer guarantee exec // if the statements are registered before the panic or return
	//return
	fmt.Println(3)
}
