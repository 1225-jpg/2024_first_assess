package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: client <number>")
		os.Exit(1)
	}
	number := os.Args[1]

	url := fmt.Sprintf("http://localhost:8080/factorial?number=%s", number)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Factorial of %s is: %s\n", number, string(body))
}