package main

import (
	"fmt"
	"github.com/username/reponame/db"
)

func main() {

	db.Open("postgres")
	fmt.Println(db.Conn)
	db.Conn = "mysql"
	fetchData(db.Conn)
}

func fetchData(conn string) {
	fmt.Println("fetching data from", conn)
}
