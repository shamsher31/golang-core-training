package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi this is first web server in %s", r.URL.Path[1:])
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":6060", nil)
}

