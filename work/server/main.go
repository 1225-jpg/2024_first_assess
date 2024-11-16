package main

import (
	"log"
	"net/http"
	"strconv"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	numStr := r.URL.Query().Get("number")
	if numStr == "" {
		http.Error(w, "No number provided", http.StatusBadRequest)
		return
    }
	num, err := strconv.Atoi(numStr)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}
	result := factorial(num)

	resultStr := strconv.Itoa(result)
	w.Write([]byte(resultStr))
}

func main() {
	http.HandleFunc("/factorial", handleRequest)
	log.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}