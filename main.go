package main

import (
	"fmt"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Starting server on PORT: 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
