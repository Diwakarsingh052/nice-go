package main

import (
	"errors"
	"fmt"
	"log"
)

var user = make(map[int]string)

// prefix your err variables with Err word

var ErrRecordNotFound error = errors.New("not found")

func main() {

	//fmt.Printf("%T\n", err)
	//fmt.Println(err)
	name, err := FetchRecord(100)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(name)
}

func FetchRecord(id int) (string, error) {
	name, ok := user[id]
	if !ok {
		return "", ErrRecordNotFound
	}
	return name, nil
}
