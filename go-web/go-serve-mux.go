package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {

		if req.URL.Path != "/" {
			http.NotFound(rw, req)
			return
		}
		fmt.Fprintf(rw, "Welcome to HTTP Request Multiplexer")
	})

	log.Printf("Server is running on http://localhost:6070")

	http.ListenAndServe(":6070", nil)

}
