package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Starting server on http://0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}
