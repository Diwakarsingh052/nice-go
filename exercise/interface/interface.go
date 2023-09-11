package main

import (
	"context"
	"excercise/interface/data"
	"excercise/interface/data/stores/postgres"

	"log"
)

func main() {
	//final outcome:- call the Create method of Conn struct using Storer interface

	// call newConn
	pConn := postgres.NewConn(nil)

	// call newStore
	store := data.NewStore(pConn)
	// call create method using the newStore variable
	u := data.User{
		Name:  "diwakar",
		Email: "diwakar@email.com",
	}
	err := store.Create(context.Background(), u)
	if err != nil {
		log.Fatal(err)
	}

	//ctx : context.Background() , sql.DB : nil

	//data.StorerInterface.Create()
}
