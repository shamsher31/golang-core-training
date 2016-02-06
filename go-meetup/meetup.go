package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Group struct {
	Name    string
	URL     string
	Member  int
	City    string
	Country string
}

var groups = []Group{
	{"GoSV", "http://www.meetup.com/golangsv", 194, "San Mateo", "US"},
	{"GoSF", "http://www.meetup.com/golangsf", 1393, "San Francisco", "US"},
}

func main() {
	http.HandleFunc("/api/groups", getGroups)
	log.Println("Server running at http://localhost:8090")
	http.ListenAndServe(":8090", nil)
}

func getGroups(w http.ResponseWriter, r *http.Request) {

	res := struct {
		Groups []Group
		Errors []string
	}{
		groups,
		[]string{"Something went wrong"},
	}

	enc := json.NewEncoder(w)
	_ = enc.Encode(res)

	fmt.Println(enc)
}
