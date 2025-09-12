package handlers

import (
	"fmt"
	"net/http"
)

// HelloHandler handles HTTP requests to return a greeting message
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}