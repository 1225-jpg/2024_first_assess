package main

import (
	"fmt"
	"net/http"
)

func web(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	http.HandleFunc("/", web)
	fmt.Println("hello!!!")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("wrong!!!", err)
	}
}
