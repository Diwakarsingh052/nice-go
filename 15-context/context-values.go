package main

import (
	"context"
	"fmt"
)

// req -> mid -> handlerFunc() -> CreateUser
// anything that starts with request and dies with request could go in context
// req id, auth token,
// do not put anything that is app-specific in the context like db conn, log

type key string

func main() {
	// context is used to pass values in request lifecycle or for timeouts and cancellation

	ctx := context.Background() // it creates an empty container to put values and timeouts
	//context.TODO() if not sure about what context to use then go with it
	//The provided key must be comparable and should not be of type string or any other built-in type to avoid collisions
	//between packages using context
	//var k string = "anyKey"
	var k key = "anyKey"

	ctx = context.WithValue(ctx, k, "abc") // empty container gets updated with the new value
	fetchValue(ctx, k)

}

func fetchValue(ctx context.Context, k key) { // context should be the first param in the func
	v := ctx.Value(k)
	//if v == nil {
	//
	//}
	s, ok := v.(string) // type assertion to make sure the data is of a valid type
	if !ok {
		fmt.Println("value not there or of a different type")
		return
	}
	fmt.Println(s)
}
