package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/f", helloHandler)
	http.ListenAndServe(":8080", nil)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
