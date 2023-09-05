package main

import (
	"fmt"
)

func main() {

	var i any = 10

	i = "hello"

	//i = 1000

	var s string

	s, ok := i.(string) // use type assertion and make this line work

	if !ok {

		fmt.Println("i Is not a string")

		return

	}

	fmt.Println("s has value : ", s)
}
