package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func main() {
	var q *strconv.NumError //nil

	_, err := strconv.Atoi("abc")

	// if NumError is present in the err var or not
	if err != nil {

		if errors.As(err, &q) { // reference imp // it checks whether struct is inside the error chain or not
			fmt.Println("error exists", q.Func)
		}

		log.Println(err)
		return
	}

	fmt.Println("not")
}
