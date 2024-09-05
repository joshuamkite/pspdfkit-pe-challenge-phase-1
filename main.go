package main

import (
	"fmt"
	"net/http"
)

// Define a handler function separately from `main` so that we can test it without duplicating the code.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from PSPDFKit Engineer!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
