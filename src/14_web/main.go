package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello world</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About</h1>")
}

var PORT string = ":3000"

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server starting...", PORT)
	http.ListenAndServe(PORT, nil)
}
