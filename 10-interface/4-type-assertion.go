package main

import "fmt"

type Runner interface {
	Run()
}
type Walker interface {
	Walk()
}
type WalkRunner interface {
	Runner // embedding interface
	Walker
}

type Human struct{ name string }

type Robo struct{ name string }

func (r Robo) Run() {}

func (h Human) Walk() {
}
func (h Human) Run() {
}

func main() {
	h := Human{name: "John"}
	_ = h
	robot := Robo{name: "r1"}

	var r Runner = robot
	r = h

	v, ok := r.(Human) // type assertion // check whether a type exists in the interface or not // if it is there than that struct would be returned
	//ok would be false if the type is not present
	if !ok {
		fmt.Println("you are not human, can't call the walk method")
		return
	}
	fmt.Println("calling the walk method")
	v.Walk()

}
