package main

import (
	"log"
	"os"
)

//https://www.devdungeon.com/content/working-files-go
//  read-> 4   // write ->2 // x ->1

func main() {
	f, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()

}
