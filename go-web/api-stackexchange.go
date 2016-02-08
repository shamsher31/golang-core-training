package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	var data struct {
		Items []struct {
			Name  string
			Count int
		}
	}

	res, _ := http.Get("https://api.stackexchange.com/2.2/tags?page=1&pagesize=20&order=desc&sort=popular&site=stackoverflow")
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	dec.Decode(&data)

	for _, item := range data.Items {
		fmt.Println(item.Name, item.Count)
	}
}
