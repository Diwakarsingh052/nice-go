package main

import (
	"fmt"
	"runtime/debug"
)

var msg any

func main() {

	//case we want intentional panic to happen
	//db, err := sql.Open("postgres", "pgx")
	//if err != nil {
	//	panic(err)
	//}
	//_ = db
	a := abc()
	if msg != nil {
		fmt.Println("panic msg from the main", msg)
	}
	fmt.Println("end of the main", a)

}

func abc() int {
	//recover func // it recovers from the panic and stop panic further propagation
	defer recoverFunc()
	var i []int
	i[10] = 100
	return 10
	//if ops goes wrong the func would panic
	panic("some problem")
}

func recoverFunc() {
	msg = recover() // nil means no panic // otherwise r would be the msg of the panic
	if msg != nil {
		fmt.Println("recovered from the panic", msg)
		fmt.Println(string(debug.Stack()))
	}
}
