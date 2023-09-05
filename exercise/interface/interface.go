package main

import (
	"github.com/username/reponame/exercise/interface/data"
)

func main() {
	//final outcome:- call the Create method of Conn struct using Storer interface

	// call newConn
	// call newStore
	// call create method using the newStore variable

	//ctx : context.Background() , sql.DB : nil

	data.StorerInterface.Create()
}
