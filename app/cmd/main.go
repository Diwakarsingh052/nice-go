package main

import (
	"net/http"
)

// /data?id=2

//localhost:8080/user?user_id=2

func main() {
	// call the GetUser Func
	http.HandleFunc("/user")
	http.ListenAndServe(":8080", nil)

}
