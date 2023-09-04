package main

import (
	"fmt"
	"github.com/username/reponame/sum" //moduleName/packageName
)

func main() {
	fmt.Println()
	sum.Add(2, 3)

	//fmt.Sprintf()  // look for design patterns for unexported functions
	// create a package named as greet // create an exported func name as hello
	// hello(name) -> hello name
}
