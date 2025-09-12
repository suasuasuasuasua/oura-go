package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)

	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	fmt.Println("Visit http://localhost:8080 to see Hello World!")

	log.Fatal(http.ListenAndServe(port, nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}