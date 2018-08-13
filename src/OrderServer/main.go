// OrderServer project main.go
package main

import (
	"fmt"
	"net/http"
)

func OrderFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("receive request %s from %s", r.RequestURI, r.RemoteAddr)
	w.Write([]byte("hello world"))
}

func main() {
	fmt.Println("Hello World!")
	http.HandleFunc("/order", OrderFunc)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("start server err: " + err.Error())
	}
}
