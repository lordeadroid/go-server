package main

import (
	"fmt"
	"handlers/handlers"
	"net/http"
)

func main() {
	PORT := 8000
	mux := http.NewServeMux()

	mux.HandleFunc("/next-pos", handlers.GetNextPosition)

	fmt.Println("Server started listing on PORT:", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux)
}
