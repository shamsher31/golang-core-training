package main

import (
	"fmt"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi this is first web server in %s", r.URL.Path[:])
}

func main() {
	http.HandleFunc("/", root)
	log.Printf("Server is running on http://localhost:6060")
	http.ListenAndServe(":6060", nil)
}
