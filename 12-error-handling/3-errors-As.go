package main

import (
	"errors"
	"fmt"
)

// QueryError is dedicated to work with error handling, it should not be used to work with normal data
// always user error as postfix for structure that we are using for error handling
type QueryError struct {
	Func  string
	Input string
	Err   error
}

var ErrNotFound = errors.New("not found")
var ErrMismatch = errors.New("mismatch")

func (q *QueryError) Error() string {
	return "main." + q.Func + ": " + "input " + q.Input + " " + q.Err.Error()
}
func main() {
	//_, err := strconv.Atoi("abc")
	//fmt.Println(err)
	//_, err = strconv.Atoi("xyz")
	//fmt.Println(err)
	//_, err = strconv.ParseInt("hello", 10, 64)
	//fmt.Println(err)
	//os.PathError{}
	//os.SyscallError{}

	err := SearchSomething("data")
	fmt.Println(err)
	err = SearchName("raj")
	fmt.Println(err)

	var q *QueryError       //nil
	if errors.As(err, &q) { // reference imp // it checks whether struct is inside the error chain or not
		fmt.Println("true", q.Func)
		return
	}

	fmt.Println("not")

}

func SearchSomething(s string) error {

	//do searching and if that is not found then return the error below
	// QueryError struct can be returned as an error value because error method is implemented over it.
	return &QueryError{
		Func:  "SearchSomething",
		Input: s,
		Err:   ErrNotFound,
	}

}

func SearchName(name string) error {
	//do searching and if that is not found then return the error below
	return &QueryError{
		Func:  "SearchName",
		Input: name,
		Err:   ErrMismatch,
	}
}
